package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListDORAFailuresCmd = &cobra.Command{
	Use:   "listdorafailures",
	Short: "Get a list of failure events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/dora/failures")
		fmt.Println("OperationID: ListDORAFailures")
	},
}

func init() {
	Cmd.AddCommand(ListDORAFailuresCmd)
}
