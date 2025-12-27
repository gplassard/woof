package spans

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateSpansCmd = &cobra.Command{
	Use:   "aggregate_spans",
	Short: "Aggregate spans",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansApi(client.NewAPIClient())
		res, _, err := api.AggregateSpans(client.NewContext(apiKey, appKey, site), datadogV2.SpansAggregateRequest{})
		if err != nil {
			log.Fatalf("failed to aggregate_spans: %v", err)
		}

		cmdutil.PrintJSON(res, "spans")
	},
}

func init() {
	Cmd.AddCommand(AggregateSpansCmd)
}
