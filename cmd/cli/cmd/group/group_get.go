package datumgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing new datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.Flags().StringP("id", "i", "", "group id to query")
	datum.ViperBindFlag("group.get.id", groupGetCmd.Flags().Lookup("id"))

	groupGetCmd.Flags().StringP("owner", "o", "", "get groups by owner")
	datum.ViperBindFlag("group.get.owner", groupGetCmd.Flags().Lookup("owner"))
}

func getGroup(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	gID := viper.GetString("group.get.id")
	ownerID := viper.GetString("group.get.owner")

	// if an group ID is provided, filter on that group, otherwise get all by owner
	if gID != "" {
		groups, err := cli.Client.GetGroupByID(ctx, gID, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err := json.Marshal(groups)
		if err != nil {
			return err
		}

		return datum.JSONPrint(s)
	}

	if ownerID != "" {
		whereInput := &datumclient.GroupWhereInput{
			HasOwnerWith: []*datumclient.OrganizationWhereInput{
				{
					ID: &ownerID,
				},
			},
		}

		groups, err := cli.Client.GroupsWhere(ctx, whereInput, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err := json.Marshal(groups)
		if err != nil {
			return err
		}

		return datum.JSONPrint(s)
	}

	groups, err := cli.Client.GetAllGroups(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err := json.Marshal(groups)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
