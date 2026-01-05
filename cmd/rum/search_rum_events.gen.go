package rum

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchRUMEventsCmd = &cobra.Command{
	Use:     "search-rum-events",
	Aliases: []string{"search-events"},
	Short:   "Search RUM events",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.SearchRUMEvents(client.NewContext(apiKey, appKey, site), datadogV2.RUMSearchEventsRequest{})
		cmdutil.HandleError(err, "failed to search-rum-events")

		cmd.Println(cmdutil.FormatJSON(res, "rum"))
	},
}

func init() {
	Cmd.AddCommand(SearchRUMEventsCmd)
}
