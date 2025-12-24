package logs_custom_destinations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteLogsCustomDestinationCmd = &cobra.Command{
	Use:   "deletelogscustomdestination",
	Short: "Delete a custom destination",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/logs/config/custom-destinations/{custom_destination_id}")
		fmt.Println("OperationID: DeleteLogsCustomDestination")
	},
}

func init() {
	Cmd.AddCommand(DeleteLogsCustomDestinationCmd)
}
