package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentCmd = &cobra.Command{
	Use:   "deleteincident",
	Short: "Delete an existing incident",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/{incident_id}")
		fmt.Println("OperationID: DeleteIncident")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentCmd)
}
