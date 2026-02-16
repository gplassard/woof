package events

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchEventsCmd = &cobra.Command{
	Use:     "search-events",
	Aliases: []string{"search"},
	Short:   "Search events",
	Long: `Search events
Documentation: https://docs.datadoghq.com/api/latest/events/#search-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EventsListResponse
		var err error

		optionalParams := datadogV2.NewSearchEventsOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		api := datadogV2.NewEventsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchEvents(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to search-events")

		fmt.Println(cmdutil.FormatJSON(res, "event"))
	},
}

func init() {

	SearchEventsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchEventsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchEventsCmd)
}
