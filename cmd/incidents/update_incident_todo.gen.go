package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIncidentTodoCmd = &cobra.Command{
	Use:     "update-incident-todo [incident_id] [todo_id] [payload]",
	Aliases: []string{"update-todo"},
	Short:   "Update an incident todo",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTodoResponse
		var err error

		var body datadogV2.IncidentTodoPatchRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.UpdateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-incident-todo")

		cmd.Println(cmdutil.FormatJSON(res, "incident_todos"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTodoCmd)
}
