package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAssetsSBOMsCmd = &cobra.Command{
	Use:   "list-assets-s-b-o-ms",
	Short: "List assets SBOMs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListAssetsSBOMs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-assets-s-b-o-ms: %v", err)
		}

		cmdutil.PrintJSON(res, "sboms")
	},
}

func init() {
	Cmd.AddCommand(ListAssetsSBOMsCmd)
}
