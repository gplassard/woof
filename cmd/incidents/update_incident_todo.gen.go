package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentTodoCmd = &cobra.Command{
	Use:     "update-incident-todo [incident_id] [todo_id]",
	Aliases: []string{"update-todo"},
	Short:   "Update an incident todo",
	Long: `Update an incident todo
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-todo`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTodoResponse
		var err error

		var body datadogV2.IncidentTodoPatchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-incident-todo")

		fmt.Println(cmdutil.FormatJSON(res, "incident_todos"))
	},
}

func init() {

	UpdateIncidentTodoCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentTodoCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentTodoCmd)
}
