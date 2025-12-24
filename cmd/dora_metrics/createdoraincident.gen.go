package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDORAIncidentCmd = &cobra.Command{
	Use:   "createdoraincident",
	Short: "Send an incident event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/dora/incident")
		fmt.Println("OperationID: CreateDORAIncident")
	},
}

func init() {
	Cmd.AddCommand(CreateDORAIncidentCmd)
}
