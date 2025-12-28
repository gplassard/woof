package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIncidentTypeCmd = &cobra.Command{
	Use:   "update_incident_type [incident_type_id]",
	Short: "Update an incident type",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentType(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentTypePatchRequest{})
		if err != nil {
			log.Fatalf("failed to update_incident_type: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_types")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTypeCmd)
}
