package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentTodoCmd = &cobra.Command{
	Use:     "create-incident-todo [incident_id]",
	Aliases: []string{"create-todo"},
	Short:   "Create an incident todo",
	Long: `Create an incident todo
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-todo`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTodoResponse
		var err error

		var body datadogV2.IncidentTodoCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentTodo(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-incident-todo")

		fmt.Println(cmdutil.FormatJSON(res, "incident_todo"))
	},
}

func init() {

	CreateIncidentTodoCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentTodoCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentTodoCmd)
}
