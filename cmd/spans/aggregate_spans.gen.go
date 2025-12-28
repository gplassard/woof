package spans

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateSpansCmd = &cobra.Command{
	Use:   "aggregate-spans",
	Aliases: []string{ "aggregate", },
	Short: "Aggregate spans",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansApi(client.NewAPIClient())
		res, _, err := api.AggregateSpans(client.NewContext(apiKey, appKey, site), datadogV2.SpansAggregateRequest{})
		cmdutil.HandleError(err, "failed to aggregate-spans")

		cmdutil.PrintJSON(res, "bucket")
	},
}

func init() {
	Cmd.AddCommand(AggregateSpansCmd)
}
