package rum

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRUMEventsCmd = &cobra.Command{
	Use:     "list-rum-events",
	Aliases: []string{"list-events"},
	Short:   "Get a list of RUM events",
	Long: `Get a list of RUM events
Documentation: https://docs.datadoghq.com/api/latest/rum/#list-rum-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMEventsResponse
		var err error

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRUMEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-rum-events")

		fmt.Println(cmdutil.FormatJSON(res, "rum"))
	},
}

func init() {

	Cmd.AddCommand(ListRUMEventsCmd)
}
