package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DetachCaseCmd = &cobra.Command{
	Use:   "detachcase",
	Short: "Detach security findings from their case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/security/findings/cases")
		fmt.Println("OperationID: DetachCase")
	},
}

func init() {
	Cmd.AddCommand(DetachCaseCmd)
}
