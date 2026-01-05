package rum

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AggregateRUMEventsCmd = &cobra.Command{
	Use:     "aggregate-rum-events [payload]",
	Aliases: []string{"aggregate-events"},
	Short:   "Aggregate RUM events",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMAnalyticsAggregateResponse
		var err error

		var body datadogV2.RUMAggregateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err = api.AggregateRUMEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-rum-events")

		cmd.Println(cmdutil.FormatJSON(res, "rum"))
	},
}

func init() {
	Cmd.AddCommand(AggregateRUMEventsCmd)
}
