package incidents

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchIncidentsCmd = &cobra.Command{
	Use:     "search-incidents [query]",
	Aliases: []string{"search"},
	Short:   "Search for incidents",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.SearchIncidents(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to search-incidents")

		cmd.Println(cmdutil.FormatJSON(res, "incidents_search_results"))
	},
}

func init() {
	Cmd.AddCommand(SearchIncidentsCmd)
}
