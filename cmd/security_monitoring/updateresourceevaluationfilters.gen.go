package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateResourceEvaluationFiltersCmd = &cobra.Command{
	Use:   "updateresourceevaluationfilters",
	Short: "Update resource filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/cloud_security_management/resource_filters")
		fmt.Println("OperationID: UpdateResourceEvaluationFilters")
	},
}

func init() {
	Cmd.AddCommand(UpdateResourceEvaluationFiltersCmd)
}
