package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentIntegrationCmd = &cobra.Command{
	Use:     "update-incident-integration [incident_id] [integration_metadata_id]",
	Aliases: []string{"update-integration"},
	Short:   "Update an existing incident integration metadata",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.IncidentIntegrationMetadataPatchRequest{})
		cmdutil.HandleError(err, "failed to update-incident-integration")

		cmd.Println(cmdutil.FormatJSON(res, "incident_integrations"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentIntegrationCmd)
}
