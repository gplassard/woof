package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryTimeseriesDataCmd = &cobra.Command{
	Use:   "query-timeseries-data",
	
	Short: "Query timeseries data across multiple products",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.QueryTimeseriesData(client.NewContext(apiKey, appKey, site), datadogV2.TimeseriesFormulaQueryRequest{})
		if err != nil {
			log.Fatalf("failed to query-timeseries-data: %v", err)
		}

		cmdutil.PrintJSON(res, "timeseries_response")
	},
}

func init() {
	Cmd.AddCommand(QueryTimeseriesDataCmd)
}
