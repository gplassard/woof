package processes

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListProcessesCmd = &cobra.Command{
	Use:   "listprocesses",
	Short: "Get all processes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/processes")
		fmt.Println("OperationID: ListProcesses")
	},
}

func init() {
	Cmd.AddCommand(ListProcessesCmd)
}
