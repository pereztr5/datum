package datumuser

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return users(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)

	userGetCmd.Flags().BoolP("self", "s", false, "get current user info, requires authentication")
	datum.ViperBindFlag("user.get.self", userGetCmd.Flags().Lookup("self"))

	userGetCmd.Flags().StringP("id", "i", "", "user id to query")
	datum.ViperBindFlag("user.get.id", userGetCmd.Flags().Lookup("id"))
}

func users(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	userID := viper.GetString("user.get.id")
	self := viper.GetBool("user.get.self")

	var s []byte

	writer := tables.NewTableWriter(userCmd.OutOrStdout(), "ID", "Email", "FirstName", "LastName", "AuthProvider")

	if self {
		claims, err := tokens.ParseUnverifiedTokenClaims(cli.AccessToken)
		if err != nil {
			return err
		}

		userID = claims.ParseUserID().String()
	}

	// if a user ID is provided, filter on that user, otherwise get all
	if userID != "" {
		user, err := cli.Client.GetUserByID(ctx, userID, cli.Interceptor)
		if err != nil {
			return err
		}

		if viper.GetString("output.format") == "json" {
			s, err := json.Marshal(user.User)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(user.User.ID, user.User.Email, user.User.FirstName, user.User.LastName, user.User.AuthProvider)

		writer.Render()

		return nil
	}

	users, err := cli.Client.GetAllUsers(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(users)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		return datum.JSONPrint(s)
	}

	for _, u := range users.Users.Edges {
		writer.AddRow(u.Node.ID, u.Node.Email, u.Node.FirstName, u.Node.LastName, u.Node.AuthProvider)
	}

	writer.Render()

	return nil
}
