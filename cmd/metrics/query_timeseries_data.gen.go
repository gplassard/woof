package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var QueryTimeseriesDataCmd = &cobra.Command{
	Use: "query-timeseries-data [payload]",

	Short: "Query timeseries data across multiple products",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TimeseriesFormulaQueryResponse
		var err error

		var body datadogV2.TimeseriesFormulaQueryRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err = api.QueryTimeseriesData(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-timeseries-data")

		cmd.Println(cmdutil.FormatJSON(res, "timeseries_response"))
	},
}

func init() {
	Cmd.AddCommand(QueryTimeseriesDataCmd)
}
