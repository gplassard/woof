package rum

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateRUMEventsCmd = &cobra.Command{
	Use:   "aggregate-rum-events",
	Aliases: []string{ "aggregate-events", },
	Short: "Aggregate RUM events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.AggregateRUMEvents(client.NewContext(apiKey, appKey, site), datadogV2.RUMAggregateRequest{})
		cmdutil.HandleError(err, "failed to aggregate-rum-events")

		cmdutil.PrintJSON(res, "rum")
	},
}

func init() {
	Cmd.AddCommand(AggregateRUMEventsCmd)
}
