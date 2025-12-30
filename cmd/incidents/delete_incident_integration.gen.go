package incidents

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteIncidentIntegrationCmd = &cobra.Command{
	Use:     "delete-incident-integration [incident_id] [integration_metadata_id]",
	Aliases: []string{"delete-integration"},
	Short:   "Delete an incident integration metadata",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		_, err := api.DeleteIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-incident-integration")

	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentIntegrationCmd)
}
