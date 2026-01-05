package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListUserNotificationChannelsCmd = &cobra.Command{
	Use: "list-user-notification-channels [user_id]",

	Short: "List On-Call notification channels for a user",
	Long: `List On-Call notification channels for a user
Documentation: https://docs.datadoghq.com/api/latest/on-call/#list-user-notification-channels`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListNotificationChannelsResponse
		var err error

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.ListUserNotificationChannels(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-user-notification-channels")

		cmd.Println(cmdutil.FormatJSON(res, "notification_channels"))
	},
}

func init() {

	Cmd.AddCommand(ListUserNotificationChannelsCmd)
}
