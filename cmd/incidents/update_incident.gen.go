package incidents

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentCmd = &cobra.Command{
	Use:     "update-incident [incident_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing incident",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncident(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-incident")

		cmd.Println(cmdutil.FormatJSON(res, "incidents"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentCmd)
}
