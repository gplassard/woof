package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteUserNotificationChannelCmd = &cobra.Command{
	Use:   "delete-user-notification-channel [user_id] [channel_id]",
	
	Short: "Delete an On-Call notification channel for a user",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		_, err := api.DeleteUserNotificationChannel(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to delete-user-notification-channel: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteUserNotificationChannelCmd)
}
