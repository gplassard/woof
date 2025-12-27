package incident_services

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIncidentServiceCmd = &cobra.Command{
	Use:   "delete_incident_service [service_id]",
	Short: "Delete an existing incident service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		_, err := api.DeleteIncidentService(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_incident_service: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentServiceCmd)
}
