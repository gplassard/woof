package apm_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateApmRetentionFilterCmd = &cobra.Command{
	Use:   "updateapmretentionfilter",
	Short: "Update a retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/apm/config/retention-filters/{filter_id}")
		fmt.Println("OperationID: UpdateApmRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(UpdateApmRetentionFilterCmd)
}
