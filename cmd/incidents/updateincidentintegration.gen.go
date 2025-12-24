package incidents

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIncidentIntegrationCmd = &cobra.Command{
	Use:   "updateincidentintegration [incident_id] [integration_metadata_id]",
	Short: "Update an existing incident integration metadata",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentIntegration(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.IncidentIntegrationMetadataPatchRequest{})
		if err != nil {
			log.Fatalf("failed to updateincidentintegration: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentIntegrationCmd)
}
