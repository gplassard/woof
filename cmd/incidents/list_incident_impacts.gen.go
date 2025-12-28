package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListIncidentImpactsCmd = &cobra.Command{
	Use:   "list-incident-impacts [incident_id]",
	Short: "List an incident's impacts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentImpacts(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-incident-impacts: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_impacts")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentImpactsCmd)
}
