package incident_services

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentServiceCmd = &cobra.Command{
	Use:   "getincidentservice",
	Short: "Get details of an incident service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/services/{service_id}")
		fmt.Println("OperationID: GetIncidentService")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentServiceCmd)
}
