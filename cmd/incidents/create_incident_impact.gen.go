package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentImpactCmd = &cobra.Command{
	Use:     "create-incident-impact [incident_id] [payload]",
	Aliases: []string{"create-impact"},
	Short:   "Create an incident impact",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.IncidentImpactCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentImpact(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-incident-impact")

		cmd.Println(cmdutil.FormatJSON(res, "incident_impacts"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentImpactCmd)
}
