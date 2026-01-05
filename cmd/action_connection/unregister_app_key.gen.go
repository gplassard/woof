package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UnregisterAppKeyCmd = &cobra.Command{
	Use: "unregister-app-key [app_key_id]",

	Short: "Unregister an App Key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		_, err = api.UnregisterAppKey(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to unregister-app-key")

	},
}

func init() {
	Cmd.AddCommand(UnregisterAppKeyCmd)
}
