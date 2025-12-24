package incident_services

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentServiceCmd = &cobra.Command{
	Use:   "updateincidentservice",
	Short: "Update an existing incident service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/services/{service_id}")
		fmt.Println("OperationID: UpdateIncidentService")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentServiceCmd)
}
