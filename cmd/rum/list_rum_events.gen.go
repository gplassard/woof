package rum

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRUMEventsCmd = &cobra.Command{
	Use:     "list-rum-events",
	Aliases: []string{"list-events"},
	Short:   "Get a list of RUM events",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.ListRUMEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-rum-events")

		cmd.Println(cmdutil.FormatJSON(res, "rum"))
	},
}

func init() {
	Cmd.AddCommand(ListRUMEventsCmd)
}
