package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentTodoCmd = &cobra.Command{
	Use:   "create-incident-todo [incident_id]",
	Short: "Create an incident todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentTodoCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-incident-todo: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_todos")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTodoCmd)
}
