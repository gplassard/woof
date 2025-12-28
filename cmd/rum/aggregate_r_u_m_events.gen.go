package rum

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateRUMEventsCmd = &cobra.Command{
	Use:   "aggregate-r-u-m-events",
	
	Short: "Aggregate RUM events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.AggregateRUMEvents(client.NewContext(apiKey, appKey, site), datadogV2.RUMAggregateRequest{})
		if err != nil {
			log.Fatalf("failed to aggregate-r-u-m-events: %v", err)
		}

		cmdutil.PrintJSON(res, "rum")
	},
}

func init() {
	Cmd.AddCommand(AggregateRUMEventsCmd)
}
