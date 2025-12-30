package incident_services

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteIncidentServiceCmd = &cobra.Command{
	Use:     "delete-incident-service [service_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an existing incident service",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		_, err := api.DeleteIncidentService(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-incident-service")

	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentServiceCmd)
}
