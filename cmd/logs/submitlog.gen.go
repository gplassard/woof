package logs

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SubmitLogCmd = &cobra.Command{
	Use:   "submitlog",
	Short: "Send logs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs")
		fmt.Println("OperationID: SubmitLog")
	},
}

func init() {
	Cmd.AddCommand(SubmitLogCmd)
}
