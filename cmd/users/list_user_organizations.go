package users

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/spf13/cobra"
)

var listUserOrganizations = &cobra.Command{
	Use:   "organizations [user_id]",
	Short: "Get a user organization",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userID := args[0]
		apiKey, appKey, site := util.GetConfig()
		usersApi := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := usersApi.ListUserOrganizations(client.NewContext(apiKey, appKey, site), userID)
		if err != nil {
			log.Fatalf("failed to list user organizations: %v", err)
		}

		cmdutil.PrintJSON(res, "users")
	},
}
