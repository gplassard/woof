package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDORAFailureCmd = &cobra.Command{
	Use:   "getdorafailure",
	Short: "Get a failure event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/dora/failures/{failure_id}")
		fmt.Println("OperationID: GetDORAFailure")
	},
}

func init() {
	Cmd.AddCommand(GetDORAFailureCmd)
}
