package incidents

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentIntegrationCmd = &cobra.Command{
	Use:   "createincidentintegration [incident_id]",
	Short: "Create an incident integration metadata",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentIntegrationMetadataCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createincidentintegration: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentIntegrationCmd)
}
