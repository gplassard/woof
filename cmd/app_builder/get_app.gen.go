package app_builder

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var GetAppCmd = &cobra.Command{
	Use: "get-app [app_id]",

	Short: "Get App",
	Long: `Get App
Documentation: https://docs.datadoghq.com/api/latest/app-builder/#get-app`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetAppResponse
		var err error

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetApp(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-app")

		cmd.Println(cmdutil.FormatJSON(res, "appDefinitions"))
	},
}

func init() {

	Cmd.AddCommand(GetAppCmd)
}
