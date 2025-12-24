package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentIntegrationCmd = &cobra.Command{
	Use:   "createincidentintegration",
	Short: "Create an incident integration metadata",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents/{incident_id}/relationships/integrations")
		fmt.Println("OperationID: CreateIncidentIntegration")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentIntegrationCmd)
}
