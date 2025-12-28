package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIncidentTypeCmd = &cobra.Command{
	Use:   "delete-incident-type [incident_type_id]",
	Aliases: []string{ "delete-type", },
	Short: "Delete an incident type",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		_, err := api.DeleteIncidentType(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-incident-type: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentTypeCmd)
}
