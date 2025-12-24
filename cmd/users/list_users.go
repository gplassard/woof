package users

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/spf13/cobra"
)

var listUsers = &cobra.Command{
	Use:   "list",
	Short: "List Datadog users",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		usersApi := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := usersApi.ListUsers(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list users: %v", err)
		}

		cmdutil.PrintJSON(res, "users")
	},
}
