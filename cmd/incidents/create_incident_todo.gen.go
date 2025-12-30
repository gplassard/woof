package incidents

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentTodoCmd = &cobra.Command{
	Use:     "create-incident-todo [incident_id]",
	Aliases: []string{"create-todo"},
	Short:   "Create an incident todo",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentTodoCreateRequest{})
		cmdutil.HandleError(err, "failed to create-incident-todo")

		cmd.Println(cmdutil.FormatJSON(res, "incident_todos"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTodoCmd)
}
