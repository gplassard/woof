package rum

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AggregateRUMEventsCmd = &cobra.Command{
	Use:     "aggregate-rum-events",
	Aliases: []string{"aggregate-events"},
	Short:   "Aggregate RUM events",
	Long: `Aggregate RUM events
Documentation: https://docs.datadoghq.com/api/latest/rum/#aggregate-rum-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMAnalyticsAggregateResponse
		var err error

		var body datadogV2.RUMAggregateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AggregateRUMEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-rum-events")

		cmd.Println(cmdutil.FormatJSON(res, "rum"))
	},
}

func init() {

	AggregateRUMEventsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AggregateRUMEventsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AggregateRUMEventsCmd)
}
