package apm_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetApmRetentionFilterCmd = &cobra.Command{
	Use:   "getapmretentionfilter",
	Short: "Get a given APM retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/apm/config/retention-filters/{filter_id}")
		fmt.Println("OperationID: GetApmRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(GetApmRetentionFilterCmd)
}
