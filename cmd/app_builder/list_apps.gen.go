package app_builder

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAppsCmd = &cobra.Command{
	Use: "list-apps",

	Short: "List Apps",
	Long: `List Apps
Documentation: https://docs.datadoghq.com/api/latest/app-builder/#list-apps`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListAppsResponse
		var err error

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListApps(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-apps")

		fmt.Println(cmdutil.FormatJSON(res, "appDefinitions"))
	},
}

func init() {

	Cmd.AddCommand(ListAppsCmd)
}
