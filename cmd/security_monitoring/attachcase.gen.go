package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AttachCaseCmd = &cobra.Command{
	Use:   "attachcase",
	Short: "Attach security findings to a case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security/findings/cases/{case_id}")
		fmt.Println("OperationID: AttachCase")
	},
}

func init() {
	Cmd.AddCommand(AttachCaseCmd)
}
