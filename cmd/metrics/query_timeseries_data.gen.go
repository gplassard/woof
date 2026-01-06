package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryTimeseriesDataCmd = &cobra.Command{
	Use: "query-timeseries-data",

	Short: "Query timeseries data across multiple products",
	Long: `Query timeseries data across multiple products
Documentation: https://docs.datadoghq.com/api/latest/metrics/#query-timeseries-data`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TimeseriesFormulaQueryResponse
		var err error

		var body datadogV2.TimeseriesFormulaQueryRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.QueryTimeseriesData(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-timeseries-data")

		cmd.Println(cmdutil.FormatJSON(res, "timeseries_response"))
	},
}

func init() {

	QueryTimeseriesDataCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	QueryTimeseriesDataCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(QueryTimeseriesDataCmd)
}
