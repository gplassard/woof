package logs_custom_destinations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateLogsCustomDestinationCmd = &cobra.Command{
	Use:   "updatelogscustomdestination",
	Short: "Update a custom destination",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/logs/config/custom-destinations/{custom_destination_id}")
		fmt.Println("OperationID: UpdateLogsCustomDestination")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsCustomDestinationCmd)
}
