package cloud_network_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAggregatedConnectionsCmd = &cobra.Command{
	Use: "get-aggregated-connections",

	Short: "Get all aggregated connections",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewCloudNetworkMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetAggregatedConnections(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-aggregated-connections")

		cmd.Println(cmdutil.FormatJSON(res, "aggregated_connection"))
	},
}

func init() {
	Cmd.AddCommand(GetAggregatedConnectionsCmd)
}
