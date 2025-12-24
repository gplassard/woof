package rum_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateRetentionFilterCmd = &cobra.Command{
	Use:   "createretentionfilter",
	Short: "Create a RUM retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/rum/applications/{app_id}/retention_filters")
		fmt.Println("OperationID: CreateRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(CreateRetentionFilterCmd)
}
