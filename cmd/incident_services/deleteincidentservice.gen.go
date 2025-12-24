package incident_services

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentServiceCmd = &cobra.Command{
	Use:   "deleteincidentservice",
	Short: "Delete an existing incident service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/services/{service_id}")
		fmt.Println("OperationID: DeleteIncidentService")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentServiceCmd)
}
