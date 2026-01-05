package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentIntegrationCmd = &cobra.Command{
	Use:     "create-incident-integration [incident_id] [payload]",
	Aliases: []string{"create-integration"},
	Short:   "Create an incident integration metadata",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.IncidentIntegrationMetadataCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-incident-integration")

		cmd.Println(cmdutil.FormatJSON(res, "incident_integrations"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentIntegrationCmd)
}
