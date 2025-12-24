package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentIntegrationCmd = &cobra.Command{
	Use:   "updateincidentintegration",
	Short: "Update an existing incident integration metadata",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/incidents/{incident_id}/relationships/integrations/{integration_metadata_id}")
		fmt.Println("OperationID: UpdateIncidentIntegration")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentIntegrationCmd)
}
