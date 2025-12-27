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
	Use:   "get_s_b_o_m [asset_type] [filter[asset_name]]",
	Short: "Get SBOM",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSBOM(client.NewContext(apiKey, appKey, site), datadogV2.AssetType(args[0]), args[1])
		if err != nil {
			log.Fatalf("failed to get_s_b_o_m: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(GetSBOMCmd)
}
