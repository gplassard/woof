package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentTypeCmd = &cobra.Command{
	Use:   "createincidenttype",
	Short: "Create an incident type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents/config/types")
		fmt.Println("OperationID: CreateIncidentType")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTypeCmd)
}
