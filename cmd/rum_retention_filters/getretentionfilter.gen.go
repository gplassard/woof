package rum_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRetentionFilterCmd = &cobra.Command{
	Use:   "getretentionfilter",
	Short: "Get a RUM retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/rum/applications/{app_id}/retention_filters/{rf_id}")
		fmt.Println("OperationID: GetRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(GetRetentionFilterCmd)
}
