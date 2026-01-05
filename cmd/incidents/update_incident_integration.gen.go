package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIncidentIntegrationCmd = &cobra.Command{
	Use:     "update-incident-integration [incident_id] [integration_metadata_id] [payload]",
	Aliases: []string{"update-integration"},
	Short:   "Update an existing incident integration metadata",
	Long: `Update an existing incident integration metadata
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-integration`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentIntegrationMetadataResponse
		var err error

		var body datadogV2.IncidentIntegrationMetadataPatchRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.UpdateIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-incident-integration")

		cmd.Println(cmdutil.FormatJSON(res, "incident_integrations"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentIntegrationCmd)
}
