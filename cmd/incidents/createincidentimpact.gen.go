package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentImpactCmd = &cobra.Command{
	Use:   "createincidentimpact",
	Short: "Create an incident impact",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents/{incident_id}/impacts")
		fmt.Println("OperationID: CreateIncidentImpact")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentImpactCmd)
}
