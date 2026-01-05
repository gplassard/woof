package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentTypeCmd = &cobra.Command{
	Use:     "create-incident-type",
	Aliases: []string{"create-type"},
	Short:   "Create an incident type",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentType(client.NewContext(apiKey, appKey, site), datadogV2.IncidentTypeCreateRequest{})
		cmdutil.HandleError(err, "failed to create-incident-type")

		cmd.Println(cmdutil.FormatJSON(res, "incident_types"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTypeCmd)
}
