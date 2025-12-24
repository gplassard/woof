package logs_custom_destinations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateLogsCustomDestinationCmd = &cobra.Command{
	Use:   "createlogscustomdestination",
	Short: "Create a custom destination",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/config/custom-destinations")
		fmt.Println("OperationID: CreateLogsCustomDestination")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsCustomDestinationCmd)
}
