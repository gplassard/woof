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
	Use:   "update-incident-service [service_id]",
	Aliases: []string{ "update", },
	Short: "Update an existing incident service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentService(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentServiceUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-incident-service: %v", err)
		}

		cmdutil.PrintJSON(res, "services")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentServiceCmd)
}
