package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListGlobalIncidentHandlesCmd = &cobra.Command{
	Use:     "list-global-incident-handles",
	Aliases: []string{"list-global-handles"},
	Short:   "List global incident handles",
	Long: `List global incident handles
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-global-incident-handles`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentHandlesResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListGlobalIncidentHandles(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-global-incident-handles")

		fmt.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {

	Cmd.AddCommand(ListGlobalIncidentHandlesCmd)
}
