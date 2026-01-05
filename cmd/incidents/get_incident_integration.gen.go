package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIncidentIntegrationCmd = &cobra.Command{
	Use:     "get-incident-integration [incident_id] [integration_metadata_id]",
	Aliases: []string{"get-integration"},
	Short:   "Get incident integration metadata details",
	Long: `Get incident integration metadata details
Documentation: https://docs.datadoghq.com/api/latest/incidents/#get-incident-integration`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentIntegrationMetadataResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.GetIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-incident-integration")

		cmd.Println(cmdutil.FormatJSON(res, "incident_integrations"))
	},
}

func init() {
	Cmd.AddCommand(GetIncidentIntegrationCmd)
}
