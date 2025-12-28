package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListVulnerableAssetsCmd = &cobra.Command{
	Use:   "list_vulnerable_assets",
	Short: "List vulnerable assets",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListVulnerableAssets(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_vulnerable_assets: %v", err)
		}

		cmdutil.PrintJSON(res, "assets")
	},
}

func init() {
	Cmd.AddCommand(ListVulnerableAssetsCmd)
}
