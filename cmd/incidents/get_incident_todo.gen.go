package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetIncidentTodoCmd = &cobra.Command{
	Use:   "get-incident-todo [incident_id] [todo_id]",
	Aliases: []string{ "get-todo", },
	Short: "Get incident todo details",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-incident-todo")

		cmd.Println(cmdutil.FormatJSON(res, "incident_todos"))
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTodoCmd)
}
