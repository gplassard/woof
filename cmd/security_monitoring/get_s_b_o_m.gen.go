package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSBOMCmd = &cobra.Command{
	Use:   "get-s-b-o-m [asset_type] [filter[asset_name]]",
	Short: "Get SBOM",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSBOM(client.NewContext(apiKey, appKey, site), datadogV2.AssetType(args[0]), args[1])
		if err != nil {
			log.Fatalf("failed to get-s-b-o-m: %v", err)
		}

		cmdutil.PrintJSON(res, "sboms")
	},
}

func init() {
	Cmd.AddCommand(GetSBOMCmd)
}
