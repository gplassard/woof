package cloud_network_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAggregatedConnectionsCmd = &cobra.Command{
	Use: "get-aggregated-connections",

	Short: "Get all aggregated connections",
	Long: `Get all aggregated connections
Documentation: https://docs.datadoghq.com/api/latest/cloud-network-monitoring/#get-aggregated-connections`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SingleAggregatedConnectionResponseArray
		var err error

		api := datadogV2.NewCloudNetworkMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAggregatedConnections(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-aggregated-connections")

		fmt.Println(cmdutil.FormatJSON(res, "aggregated_connection"))
	},
}

func init() {

	Cmd.AddCommand(GetAggregatedConnectionsCmd)
}
