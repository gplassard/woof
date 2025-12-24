package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUserFacetInfoCmd = &cobra.Command{
	Use:   "getuserfacetinfo",
	Short: "Get user facet info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/product-analytics/users/facet_info")
		fmt.Println("OperationID: GetUserFacetInfo")
	},
}

func init() {
	Cmd.AddCommand(GetUserFacetInfoCmd)
}
