package rum_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateRetentionFilterCmd = &cobra.Command{
	Use:   "updateretentionfilter",
	Short: "Update a RUM retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/rum/applications/{app_id}/retention_filters/{rf_id}")
		fmt.Println("OperationID: UpdateRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(UpdateRetentionFilterCmd)
}
