package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentTypeCmd = &cobra.Command{
	Use:     "create-incident-type [payload]",
	Aliases: []string{"create-type"},
	Short:   "Create an incident type",
	Long: `Create an incident type
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-type`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTypeResponse
		var err error

		var body datadogV2.IncidentTypeCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.CreateIncidentType(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-type")

		cmd.Println(cmdutil.FormatJSON(res, "incident_types"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTypeCmd)
}
