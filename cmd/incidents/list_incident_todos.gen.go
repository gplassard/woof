package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListIncidentTodosCmd = &cobra.Command{
	Use:   "list-incident-todos [incident_id]",
	Aliases: []string{ "list-todos", },
	Short: "Get a list of an incident's todos",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentTodos(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-incident-todos: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_todos")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTodosCmd)
}
