package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateUserNotificationChannelCmd = &cobra.Command{
	Use:   "create-user-notification-channel [user_id]",
	
	Short: "Create an On-Call notification channel for a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.CreateUserNotificationChannel(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CreateUserNotificationChannelRequest{})
		if err != nil {
			log.Fatalf("failed to create-user-notification-channel: %v", err)
		}

		cmdutil.PrintJSON(res, "notification_channels")
	},
}

func init() {
	Cmd.AddCommand(CreateUserNotificationChannelCmd)
}
