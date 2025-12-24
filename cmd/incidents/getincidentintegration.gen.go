package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentIntegrationCmd = &cobra.Command{
	Use:   "getincidentintegration",
	Short: "Get incident integration metadata details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/{incident_id}/relationships/integrations/{integration_metadata_id}")
		fmt.Println("OperationID: GetIncidentIntegration")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentIntegrationCmd)
}
