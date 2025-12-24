package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentCmd = &cobra.Command{
	Use:   "updateincident",
	Short: "Update an existing incident",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/incidents/{incident_id}")
		fmt.Println("OperationID: UpdateIncident")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentCmd)
}
