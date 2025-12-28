package users

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateUserCmd = &cobra.Command{
	Use:   "update-user [user_id]",
	Short: "Update a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.UpdateUser(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UserUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-user: %v", err)
		}

		cmdutil.PrintJSON(res, "users")
	},
}

func init() {
	Cmd.AddCommand(UpdateUserCmd)
}
