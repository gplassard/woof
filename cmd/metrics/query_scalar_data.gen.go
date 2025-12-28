package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var QueryScalarDataCmd = &cobra.Command{
	Use:   "query-scalar-data",
	
	Short: "Query scalar data across multiple products",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.QueryScalarData(client.NewContext(apiKey, appKey, site), datadogV2.ScalarFormulaQueryRequest{})
		if err != nil {
			log.Fatalf("failed to query-scalar-data: %v", err)
		}

		cmdutil.PrintJSON(res, "scalar_response")
	},
}

func init() {
	Cmd.AddCommand(QueryScalarDataCmd)
}
