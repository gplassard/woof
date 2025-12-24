package incident_services

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentServiceCmd = &cobra.Command{
	Use:   "createincidentservice",
	Short: "Create a new incident service",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentServicesApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentService(client.NewContext(apiKey, appKey, site), datadogV2.IncidentServiceCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createincidentservice: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_services")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentServiceCmd)
}
