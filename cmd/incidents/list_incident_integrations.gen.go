package incidents

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentIntegrationsCmd = &cobra.Command{
	Use:     "list-incident-integrations [incident_id]",
	Aliases: []string{"list-integrations"},
	Short:   "Get a list of an incident's integration metadata",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentIntegrations(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-incident-integrations")

		cmd.Println(cmdutil.FormatJSON(res, "incident_integrations"))
	},
}

func init() {
	Cmd.AddCommand(ListIncidentIntegrationsCmd)
}
