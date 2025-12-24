package logs_custom_destinations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetLogsCustomDestinationCmd = &cobra.Command{
	Use:   "getlogscustomdestination",
	Short: "Get a custom destination",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/custom-destinations/{custom_destination_id}")
		fmt.Println("OperationID: GetLogsCustomDestination")
	},
}

func init() {
	Cmd.AddCommand(GetLogsCustomDestinationCmd)
}
