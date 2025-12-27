package cloud_network_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAggregatedDnsCmd = &cobra.Command{
	Use:   "getaggregateddns",
	Short: "Get all aggregated DNS traffic",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudNetworkMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetAggregatedDns(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getaggregateddns: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_network_monitoring")
	},
}

func init() {
	Cmd.AddCommand(GetAggregatedDnsCmd)
}
