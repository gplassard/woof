package cloud_network_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAggregatedConnectionsCmd = &cobra.Command{
	Use:   "getaggregatedconnections",
	Short: "Get all aggregated connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudNetworkMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetAggregatedConnections(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getaggregatedconnections: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_network_monitoring")
	},
}

func init() {
	Cmd.AddCommand(GetAggregatedConnectionsCmd)
}
