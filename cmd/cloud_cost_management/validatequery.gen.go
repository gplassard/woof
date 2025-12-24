package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ValidateQueryCmd = &cobra.Command{
	Use:   "validatequery",
	Short: "Validate query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/tags/enrichment/validate-query")
		fmt.Println("OperationID: ValidateQuery")
	},
}

func init() {
	Cmd.AddCommand(ValidateQueryCmd)
}
