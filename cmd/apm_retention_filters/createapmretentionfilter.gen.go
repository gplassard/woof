package apm_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateApmRetentionFilterCmd = &cobra.Command{
	Use:   "createapmretentionfilter",
	Short: "Create a retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/apm/config/retention-filters")
		fmt.Println("OperationID: CreateApmRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(CreateApmRetentionFilterCmd)
}
