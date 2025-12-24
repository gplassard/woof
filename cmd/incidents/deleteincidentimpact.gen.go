package incidents

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIncidentImpactCmd = &cobra.Command{
	Use:   "deleteincidentimpact [incident_id] [impact_id]",
	Short: "Delete an incident impact",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		_, err := api.DeleteIncidentImpact(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deleteincidentimpact: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentImpactCmd)
}
