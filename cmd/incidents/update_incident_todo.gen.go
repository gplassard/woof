package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIncidentTodoCmd = &cobra.Command{
	Use:   "update-incident-todo [incident_id] [todo_id]",
	Short: "Update an incident todo",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.IncidentTodoPatchRequest{})
		if err != nil {
			log.Fatalf("failed to update-incident-todo: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_todos")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTodoCmd)
}
