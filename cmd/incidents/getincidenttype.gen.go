package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentTypeCmd = &cobra.Command{
	Use:   "getincidenttype",
	Short: "Get incident type details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/config/types/{incident_type_id}")
		fmt.Println("OperationID: GetIncidentType")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTypeCmd)
}
