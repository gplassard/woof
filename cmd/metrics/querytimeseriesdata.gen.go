package metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryTimeseriesDataCmd = &cobra.Command{
	Use:   "querytimeseriesdata",
	Short: "Query timeseries data across multiple products",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.QueryTimeseriesData(client.NewContext(apiKey, appKey, site), datadogV2.TimeseriesFormulaQueryRequest{})
		if err != nil {
			log.Fatalf("failed to querytimeseriesdata: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(QueryTimeseriesDataCmd)
}
