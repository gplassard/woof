package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIncidentTodoCmd = &cobra.Command{
	Use:     "get-incident-todo [incident_id] [todo_id]",
	Aliases: []string{"get-todo"},
	Short:   "Get incident todo details",
	Long: `Get incident todo details
Documentation: https://docs.datadoghq.com/api/latest/incidents/#get-incident-todo`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTodoResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.GetIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-incident-todo")

		cmd.Println(cmdutil.FormatJSON(res, "incident_todos"))
	},
}

func init() {

	Cmd.AddCommand(GetIncidentTodoCmd)
}
