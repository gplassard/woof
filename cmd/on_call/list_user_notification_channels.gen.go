package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListUserNotificationChannelsCmd = &cobra.Command{
	Use:   "list_user_notification_channels [user_id]",
	Short: "List On-Call notification channels for a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.ListUserNotificationChannels(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list_user_notification_channels: %v", err)
		}

		cmdutil.PrintJSON(res, "notification_channels")
	},
}

func init() {
	Cmd.AddCommand(ListUserNotificationChannelsCmd)
}
