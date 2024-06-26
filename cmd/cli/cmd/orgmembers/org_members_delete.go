package datumorgmembers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var orgMembersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a user from a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteOrgMember(cmd.Context())
	},
}

func init() {
	orgMembersCmd.AddCommand(orgMembersDeleteCmd)

	orgMembersDeleteCmd.Flags().StringP("org-id", "o", "", "organization id")
	datum.ViperBindFlag("orgmember.delete.orgid", orgMembersDeleteCmd.Flags().Lookup("org-id"))

	orgMembersDeleteCmd.Flags().StringP("user-id", "u", "", "user id")
	datum.ViperBindFlag("orgmember.delete.userid", orgMembersDeleteCmd.Flags().Lookup("user-id"))
}

func deleteOrgMember(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("orgmember.delete.orgid")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	uID := viper.GetString("orgmember.delete.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// get the id of the org member
	where := datumclient.OrgMembershipWhereInput{
		OrganizationID: &oID,
		UserID:         &uID,
	}

	orgMembers, err := cli.Client.GetOrgMembersByOrgID(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	if len(orgMembers.OrgMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:goerr113
	}

	id := orgMembers.OrgMemberships.Edges[0].Node.ID

	var s []byte

	orgMember, err := cli.Client.RemoveUserFromOrg(ctx, id, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(orgMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
