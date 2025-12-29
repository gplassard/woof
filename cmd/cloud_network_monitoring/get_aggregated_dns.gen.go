package cloud_network_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAggregatedDnsCmd = &cobra.Command{
	Use:   "get-aggregated-dns",
	
	Short: "Get all aggregated DNS traffic",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudNetworkMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetAggregatedDns(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-aggregated-dns")

		cmd.Println(cmdutil.FormatJSON(res, "aggregated_dns"))
	},
}

func init() {
	Cmd.AddCommand(GetAggregatedDnsCmd)
}
