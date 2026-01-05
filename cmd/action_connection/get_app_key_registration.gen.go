package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAppKeyRegistrationCmd = &cobra.Command{
	Use: "get-app-key-registration [app_key_id]",

	Short: "Get an existing App Key Registration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.GetAppKeyRegistration(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-app-key-registration")

		cmd.Println(cmdutil.FormatJSON(res, "app_key_registration"))
	},
}

func init() {
	Cmd.AddCommand(GetAppKeyRegistrationCmd)
}
