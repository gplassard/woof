package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentIntegrationCmd = &cobra.Command{
	Use:   "deleteincidentintegration",
	Short: "Delete an incident integration metadata",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/{incident_id}/relationships/integrations/{integration_metadata_id}")
		fmt.Println("OperationID: DeleteIncidentIntegration")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentIntegrationCmd)
}
