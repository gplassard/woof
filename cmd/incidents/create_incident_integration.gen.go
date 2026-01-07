package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentIntegrationCmd = &cobra.Command{
	Use:     "create-incident-integration [incident_id]",
	Aliases: []string{"create-integration"},
	Short:   "Create an incident integration metadata",
	Long: `Create an incident integration metadata
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-integration`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentIntegrationMetadataResponse
		var err error

		var body datadogV2.IncidentIntegrationMetadataCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-incident-integration")

		fmt.Println(cmdutil.FormatJSON(res, "incident_integrations"))
	},
}

func init() {

	CreateIncidentIntegrationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentIntegrationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentIntegrationCmd)
}
