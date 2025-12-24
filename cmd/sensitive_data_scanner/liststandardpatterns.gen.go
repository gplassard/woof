package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListStandardPatternsCmd = &cobra.Command{
	Use:   "liststandardpatterns",
	Short: "List standard patterns",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/sensitive-data-scanner/config/standard-patterns")
		fmt.Println("OperationID: ListStandardPatterns")
	},
}

func init() {
	Cmd.AddCommand(ListStandardPatternsCmd)
}
