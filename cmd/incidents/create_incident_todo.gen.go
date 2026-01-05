package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentTodoCmd = &cobra.Command{
	Use:     "create-incident-todo [incident_id] [payload]",
	Aliases: []string{"create-todo"},
	Short:   "Create an incident todo",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTodoResponse
		var err error

		var body datadogV2.IncidentTodoCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.CreateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-incident-todo")

		cmd.Println(cmdutil.FormatJSON(res, "incident_todos"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTodoCmd)
}
