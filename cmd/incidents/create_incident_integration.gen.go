package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentIntegrationCmd = &cobra.Command{
	Use:   "create-incident-integration [incident_id]",
	Aliases: []string{ "create-integration", },
	Short: "Create an incident integration metadata",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentIntegrationMetadataCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-incident-integration: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_integrations")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentIntegrationCmd)
}
