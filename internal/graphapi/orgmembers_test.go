package graphapi_test

import (
	"fmt"
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/pkg/datumclient"
)

func (suite *GraphTestSuite) TestQueryOrgMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// allow access to organization
	checkCtx := privacy.DecisionContext(reqCtx, privacy.Allow)

	orgMember, err := org1.Members(checkCtx)
	require.NoError(t, err)
	require.Len(t, orgMember, 1)

	testCases := []struct {
		name      string
		queryID   string
		allowed   bool
		expected  *ent.OrgMembership
		expectErr bool
	}{
		{
			name:     "happy path, get org member by org id",
			queryID:  org1.ID,
			allowed:  true,
			expected: orgMember[0],
		},
		{
			name:      "no access",
			queryID:   org1.ID,
			allowed:   false,
			expected:  nil,
			expectErr: true,
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			allowed:  true,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			orgID := tc.queryID
			whereInput := datumclient.OrgMembershipWhereInput{
				OrganizationID: &orgID,
			}

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.GetOrgMembersByOrgID(reqCtx, &whereInput)

			if tc.expectErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)

			if tc.expected == nil {
				assert.Empty(t, resp.OrgMemberships.Edges)

				return
			}

			require.NotNil(t, resp)
			require.NotNil(t, resp.OrgMemberships)
			assert.Equal(t, tc.expected.UserID, resp.OrgMemberships.Edges[0].Node.UserID)
			assert.Equal(t, tc.expected.Role, resp.OrgMemberships.Edges[0].Node.Role)
		})
	}

	// delete created org
	(&OrganizationCleanup{client: suite.client, OrgID: org1.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryCreateOrgMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	personalOrg := (&OrganizationBuilder{client: suite.client, PersonalOrg: true}).MustNew(reqCtx, t)
	listObjects := []string{fmt.Sprintf("organization:%s", org1.ID), fmt.Sprintf("organization:%s", personalOrg.ID)}

	// allow access to organization
	checkCtx := privacy.DecisionContext(reqCtx, privacy.Allow)

	orgMember, err := org1.Members(checkCtx)
	require.NoError(t, err)
	require.Len(t, orgMember, 1)

	testUser1 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	testUser2 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name      string
		orgID     string
		userID    string
		role      enums.Role
		checkOrg  bool
		checkRole bool
		errMsg    string
	}{
		{
			name:      "happy path, add admin",
			orgID:     org1.ID,
			userID:    testUser1.ID,
			role:      enums.RoleAdmin,
			checkRole: true,
			checkOrg:  true,
		},
		{
			name:      "happy path, add member",
			orgID:     org1.ID,
			userID:    testUser2.ID,
			role:      enums.RoleMember,
			checkRole: true,
			checkOrg:  true,
		},
		{
			name:      "duplicate user, different role",
			orgID:     org1.ID,
			userID:    testUser1.ID,
			role:      enums.RoleMember,
			checkOrg:  true,
			checkRole: true,
			errMsg:    "constraint failed",
		},
		{
			name:      "add user to personal org not allowed",
			orgID:     personalOrg.ID,
			userID:    testUser1.ID,
			role:      enums.RoleMember,
			checkOrg:  true,
			checkRole: true,
			errMsg:    hooks.ErrPersonalOrgsNoMembers.Error(),
		},
		{
			name:      "invalid user",
			orgID:     org1.ID,
			userID:    "not-a-valid-user-id",
			role:      enums.RoleMember,
			checkOrg:  true,
			checkRole: true,
			errMsg:    "constraint failed", // TODO: better error messaging: https://github.com/datumforge/datum/issues/415
		},
		{
			name:      "invalid org",
			orgID:     "not-a-valid-org-id",
			userID:    testUser1.ID,
			role:      enums.RoleMember,
			checkOrg:  true,
			checkRole: true,
			errMsg:    "organization not found",
		},
		{
			name:      "invalid role",
			orgID:     org1.ID,
			userID:    testUser1.ID,
			role:      enums.Invalid,
			checkOrg:  false,
			checkRole: false,
			errMsg:    "not a valid OrgMembershipRole",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.errMsg == "" {
				mock_fga.WriteAny(t, suite.client.fga)
			}

			if tc.checkOrg {
				// checks for adding orgs to ensure not a personal org
				mock_fga.ListAny(t, suite.client.fga, listObjects)
			}

			// checks role in org to ensure user has ability to add other members
			if tc.checkRole {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			role := tc.role
			input := datumclient.CreateOrgMembershipInput{
				OrganizationID: tc.orgID,
				UserID:         tc.userID,
				Role:           &role,
			}

			resp, err := suite.client.datum.AddUserToOrgWithRole(reqCtx, input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateOrgMembership)
			assert.Equal(t, tc.userID, resp.CreateOrgMembership.OrgMembership.UserID)
			assert.Equal(t, tc.orgID, resp.CreateOrgMembership.OrgMembership.OrganizationID)
			assert.Equal(t, tc.role, resp.CreateOrgMembership.OrgMembership.Role)
		})
	}

	// delete created org and users
	(&OrganizationCleanup{client: suite.client, OrgID: org1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: suite.client, UserID: testUser1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: suite.client, UserID: testUser2.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryUpdateOrgMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&OrgMemberBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name       string
		role       enums.Role
		tupleWrite bool
		errMsg     string
	}{
		{
			name:       "happy path, update to admin from member",
			tupleWrite: true,
			role:       enums.RoleAdmin,
		},
		{
			name:       "happy path, update to member from admin",
			tupleWrite: true,
			role:       enums.RoleMember,
		},
		{
			name:       "update to same role",
			tupleWrite: false, // nothing should change
			role:       enums.RoleMember,
		},
		{
			name:       "invalid role",
			role:       enums.Invalid,
			tupleWrite: false,
			errMsg:     "not a valid OrgMembershipRole",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.tupleWrite {
				mock_fga.WriteAny(t, suite.client.fga)
			}

			if tc.errMsg == "" {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			role := tc.role
			input := datumclient.UpdateOrgMembershipInput{
				Role: &role,
			}

			resp, err := suite.client.datum.UpdateUserRoleInOrg(reqCtx, om.ID, input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateOrgMembership)
			assert.Equal(t, tc.role, resp.UpdateOrgMembership.OrgMembership.Role)
		})
	}

	// delete created org and users
	(&OrgMemberCleanup{client: suite.client, ID: om.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryDeleteOrgMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&OrgMemberBuilder{client: suite.client}).MustNew(reqCtx, t)

	mock_fga.WriteAny(t, suite.client.fga)
	mock_fga.CheckAny(t, suite.client.fga, true)

	resp, err := suite.client.datum.RemoveUserFromOrg(reqCtx, om.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteOrgMembership)
	assert.Equal(t, om.ID, resp.DeleteOrgMembership.DeletedID)
}
