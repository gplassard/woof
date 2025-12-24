package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDORAFailureCmd = &cobra.Command{
	Use:   "deletedorafailure",
	Short: "Delete a failure event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/dora/failure/{failure_id}")
		fmt.Println("OperationID: DeleteDORAFailure")
	},
}

func init() {
	Cmd.AddCommand(DeleteDORAFailureCmd)
}
