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

		optionalParams := datadogV2.NewGetAggregatedConnectionsOptionalParameters()

		if cmd.Flags().Changed("from") {
			val, _ := cmd.Flags().GetInt64("from")
			optionalParams.WithFrom(val)
		}

		if cmd.Flags().Changed("to") {
			val, _ := cmd.Flags().GetInt64("to")
			optionalParams.WithTo(val)
		}

		if cmd.Flags().Changed("group-by") {
			val, _ := cmd.Flags().GetString("group-by")
			optionalParams.WithGroupBy(val)
		}

		if cmd.Flags().Changed("tags") {
			val, _ := cmd.Flags().GetString("tags")
			optionalParams.WithTags(val)
		}

		if cmd.Flags().Changed("limit") {
			val, _ := cmd.Flags().GetInt64("limit")
			optionalParams.WithLimit(val)
		}

		api := datadogV2.NewCloudNetworkMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAggregatedConnections(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to get-aggregated-connections")

		fmt.Println(cmdutil.FormatJSON(res, "aggregated_connection"))
	},
}

func init() {

	GetAggregatedConnectionsCmd.Flags().Int64("from", 0, "Unix timestamp (number of seconds since epoch) of the start of the query window. If not provided, the start of the query window is 15 minutes before the 'to' timestamp. If neither 'from' nor 'to' are provided, the query window is '[now - 15m, now]'.")

	GetAggregatedConnectionsCmd.Flags().Int64("to", 0, "Unix timestamp (number of seconds since epoch) of the end of the query window. If not provided, the end of the query window is the current time. If neither 'from' nor 'to' are provided, the query window is '[now - 15m, now]'.")

	GetAggregatedConnectionsCmd.Flags().String("group-by", "", "Comma-separated list of fields to group connections by. The maximum number of group_by(s) is 10.")

	GetAggregatedConnectionsCmd.Flags().String("tags", "", "Comma-separated list of tags to filter connections by.")

	GetAggregatedConnectionsCmd.Flags().Int64("limit", 0, "The number of connections to be returned. The maximum value is 7500. The default is 100.")

	Cmd.AddCommand(GetAggregatedConnectionsCmd)
}
