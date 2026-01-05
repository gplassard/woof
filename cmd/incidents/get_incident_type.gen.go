package incidents

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIncidentTypeCmd = &cobra.Command{
	Use:     "get-incident-type [incident_type_id]",
	Aliases: []string{"get-type"},
	Short:   "Get incident type details",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentType(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-incident-type")

		cmd.Println(cmdutil.FormatJSON(res, "incident_types"))
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTypeCmd)
}
