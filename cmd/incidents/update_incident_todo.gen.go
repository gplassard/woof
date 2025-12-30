package incidents

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentTodoCmd = &cobra.Command{
	Use:     "update-incident-todo [incident_id] [todo_id]",
	Aliases: []string{"update-todo"},
	Short:   "Update an incident todo",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.IncidentTodoPatchRequest{})
		cmdutil.HandleError(err, "failed to update-incident-todo")

		cmd.Println(cmdutil.FormatJSON(res, "incident_todos"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTodoCmd)
}
