package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentTypeCmd = &cobra.Command{
	Use:   "updateincidenttype",
	Short: "Update an incident type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/incidents/config/types/{incident_type_id}")
		fmt.Println("OperationID: UpdateIncidentType")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTypeCmd)
}
