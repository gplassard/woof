package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteUserNotificationChannelCmd = &cobra.Command{
	Use: "delete-user-notification-channel [user_id] [channel_id]",

	Short: "Delete an On-Call notification channel for a user",
	Long: `Delete an On-Call notification channel for a user
Documentation: https://docs.datadoghq.com/api/latest/on-call/#delete-user-notification-channel`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		_, err = api.DeleteUserNotificationChannel(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-user-notification-channel")

	},
}

func init() {
	Cmd.AddCommand(DeleteUserNotificationChannelCmd)
}
