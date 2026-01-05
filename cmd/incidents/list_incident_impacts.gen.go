package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentImpactsCmd = &cobra.Command{
	Use:     "list-incident-impacts [incident_id]",
	Aliases: []string{"list-impacts"},
	Short:   "List an incident's impacts",
	Long: `List an incident's impacts
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-impacts`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentImpactsResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.ListIncidentImpacts(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-incident-impacts")

		cmd.Println(cmdutil.FormatJSON(res, "incident_impacts"))
	},
}

func init() {

	Cmd.AddCommand(ListIncidentImpactsCmd)
}
