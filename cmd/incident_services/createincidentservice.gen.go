package incident_services

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentServiceCmd = &cobra.Command{
	Use:   "createincidentservice",
	Short: "Create a new incident service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/services")
		fmt.Println("OperationID: CreateIncidentService")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentServiceCmd)
}
