package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteIncidentCmd = &cobra.Command{
	Use:     "delete-incident [incident_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an existing incident",
	Long: `Delete an existing incident
Documentation: https://docs.datadoghq.com/api/latest/incidents/#delete-incident`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		_, err = api.DeleteIncident(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-incident")

	},
}

func init() {

	Cmd.AddCommand(DeleteIncidentCmd)
}
