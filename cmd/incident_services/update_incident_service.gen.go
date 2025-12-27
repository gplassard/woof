package incident_services

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIncidentServiceCmd = &cobra.Command{
	Use:   "update_incident_service [service_id]",
	Short: "Update an existing incident service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentService(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentServiceUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_incident_service: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_services")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentServiceCmd)
}
