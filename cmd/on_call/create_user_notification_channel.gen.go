package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateUserNotificationChannelCmd = &cobra.Command{
	Use: "create-user-notification-channel [user_id] [payload]",

	Short: "Create an On-Call notification channel for a user",
	Long: `Create an On-Call notification channel for a user
Documentation: https://docs.datadoghq.com/api/latest/on-call/#create-user-notification-channel`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.NotificationChannel
		var err error

		var body datadogV2.CreateUserNotificationChannelRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.CreateUserNotificationChannel(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-user-notification-channel")

		cmd.Println(cmdutil.FormatJSON(res, "notification_channels"))
	},
}

func init() {
	Cmd.AddCommand(CreateUserNotificationChannelCmd)
}
