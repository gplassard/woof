package rum

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchRUMEventsCmd = &cobra.Command{
	Use:     "search-rum-events",
	Aliases: []string{"search-events"},
	Short:   "Search RUM events",
	Long: `Search RUM events
Documentation: https://docs.datadoghq.com/api/latest/rum/#search-rum-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMEventsResponse
		var err error

		var body datadogV2.RUMSearchEventsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err = api.SearchRUMEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-rum-events")

		cmd.Println(cmdutil.FormatJSON(res, "rum"))
	},
}

func init() {

	SearchRUMEventsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchRUMEventsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchRUMEventsCmd)
}
