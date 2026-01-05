package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RegisterAppKeyCmd = &cobra.Command{
	Use: "register-app-key [app_key_id]",

	Short: "Register a new App Key",
	Long: `Register a new App Key
Documentation: https://docs.datadoghq.com/api/latest/action-connection/#register-app-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RegisterAppKeyResponse
		var err error

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err = api.RegisterAppKey(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to register-app-key")

		cmd.Println(cmdutil.FormatJSON(res, "app_key_registration"))
	},
}

func init() {

	Cmd.AddCommand(RegisterAppKeyCmd)
}
