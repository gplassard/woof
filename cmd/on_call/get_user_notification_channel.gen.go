package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetUserNotificationChannelCmd = &cobra.Command{
	Use: "get-user-notification-channel [user_id] [channel_id]",

	Short: "Get an On-Call notification channel for a user",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetUserNotificationChannel(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-user-notification-channel")

		cmd.Println(cmdutil.FormatJSON(res, "notification_channels"))
	},
}

func init() {
	Cmd.AddCommand(GetUserNotificationChannelCmd)
}
