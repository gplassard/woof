package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentImpactsCmd = &cobra.Command{
	Use:   "listincidentimpacts",
	Short: "List an incident's impacts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/{incident_id}/impacts")
		fmt.Println("OperationID: ListIncidentImpacts")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentImpactsCmd)
}
