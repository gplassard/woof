package incident_services

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentServicesCmd = &cobra.Command{
	Use:   "listincidentservices",
	Short: "Get a list of all incident services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/services")
		fmt.Println("OperationID: ListIncidentServices")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentServicesCmd)
}
