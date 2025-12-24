package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDORAFailureCmd = &cobra.Command{
	Use:   "createdorafailure",
	Short: "Send a failure event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/dora/failure")
		fmt.Println("OperationID: CreateDORAFailure")
	},
}

func init() {
	Cmd.AddCommand(CreateDORAFailureCmd)
}
