package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentTypeCmd = &cobra.Command{
	Use:   "deleteincidenttype",
	Short: "Delete an incident type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/config/types/{incident_type_id}")
		fmt.Println("OperationID: DeleteIncidentType")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentTypeCmd)
}
