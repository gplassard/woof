package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentImpactCmd = &cobra.Command{
	Use:   "deleteincidentimpact",
	Short: "Delete an incident impact",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/{incident_id}/impacts/{impact_id}")
		fmt.Println("OperationID: DeleteIncidentImpact")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentImpactCmd)
}
