package logs_custom_destinations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListLogsCustomDestinationsCmd = &cobra.Command{
	Use:   "listlogscustomdestinations",
	Short: "Get all custom destinations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/custom-destinations")
		fmt.Println("OperationID: ListLogsCustomDestinations")
	},
}

func init() {
	Cmd.AddCommand(ListLogsCustomDestinationsCmd)
}
