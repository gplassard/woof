package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListScannedAssetsMetadataCmd = &cobra.Command{
	Use:   "listscannedassetsmetadata",
	Short: "List scanned assets metadata",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListScannedAssetsMetadata(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listscannedassetsmetadata: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ListScannedAssetsMetadataCmd)
}
