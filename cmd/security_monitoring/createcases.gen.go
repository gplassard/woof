package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCasesCmd = &cobra.Command{
	Use:   "createcases",
	Short: "Create cases for security findings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security/findings/cases")
		fmt.Println("OperationID: CreateCases")
	},
}

func init() {
	Cmd.AddCommand(CreateCasesCmd)
}
