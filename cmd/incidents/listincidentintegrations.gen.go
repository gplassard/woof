package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentIntegrationsCmd = &cobra.Command{
	Use:   "listincidentintegrations",
	Short: "Get a list of an incident's integration metadata",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/{incident_id}/relationships/integrations")
		fmt.Println("OperationID: ListIncidentIntegrations")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentIntegrationsCmd)
}
