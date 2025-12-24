package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetResourceEvaluationFiltersCmd = &cobra.Command{
	Use:   "getresourceevaluationfilters",
	Short: "List resource filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cloud_security_management/resource_filters")
		fmt.Println("OperationID: GetResourceEvaluationFilters")
	},
}

func init() {
	Cmd.AddCommand(GetResourceEvaluationFiltersCmd)
}
