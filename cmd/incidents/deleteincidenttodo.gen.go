package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIncidentTodoCmd = &cobra.Command{
	Use:   "deleteincidenttodo [incident_id] [todo_id]",
	Short: "Delete an incident todo",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		_, err := api.DeleteIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deleteincidenttodo: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentTodoCmd)
}
