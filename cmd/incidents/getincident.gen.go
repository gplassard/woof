package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentCmd = &cobra.Command{
	Use:   "getincident",
	Short: "Get the details of an incident",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/{incident_id}")
		fmt.Println("OperationID: GetIncident")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentCmd)
}
