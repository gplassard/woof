package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIncidentTypeCmd = &cobra.Command{
	Use:     "update-incident-type [incident_type_id] [payload]",
	Aliases: []string{"update-type"},
	Short:   "Update an incident type",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTypeResponse
		var err error

		var body datadogV2.IncidentTypePatchRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.UpdateIncidentType(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-incident-type")

		cmd.Println(cmdutil.FormatJSON(res, "incident_types"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTypeCmd)
}
