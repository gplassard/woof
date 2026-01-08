package on_call

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateUserNotificationChannelCmd = &cobra.Command{
	Use: "create-user-notification-channel [user_id]",

	Short: "Create an On-Call notification channel for a user",
	Long: `Create an On-Call notification channel for a user
Documentation: https://docs.datadoghq.com/api/latest/on-call/#create-user-notification-channel`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.NotificationChannel
		var err error

		var body datadogV2.CreateUserNotificationChannelRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateUserNotificationChannel(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-user-notification-channel")

		fmt.Println(cmdutil.FormatJSON(res, "user_notification_channel"))
	},
}

func init() {

	CreateUserNotificationChannelCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateUserNotificationChannelCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateUserNotificationChannelCmd)
}
