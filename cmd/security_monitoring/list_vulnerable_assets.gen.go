package security_monitoring

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListVulnerableAssetsCmd = &cobra.Command{
	Use: "list-vulnerable-assets",

	Short: "List vulnerable assets",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListVulnerableAssets(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-vulnerable-assets")

		cmd.Println(cmdutil.FormatJSON(res, "assets"))
	},
}

func init() {
	Cmd.AddCommand(ListVulnerableAssetsCmd)
}
