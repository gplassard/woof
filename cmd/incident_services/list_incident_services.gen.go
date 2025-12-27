package incident_services

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListIncidentServicesCmd = &cobra.Command{
	Use:   "list_incident_services",
	Short: "Get a list of all incident services",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err := api.ListIncidentServices(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_incident_services: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_services")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentServicesCmd)
}
