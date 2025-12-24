package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentCmd = &cobra.Command{
	Use:   "createincident",
	Short: "Create an incident",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents")
		fmt.Println("OperationID: CreateIncident")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentCmd)
}
