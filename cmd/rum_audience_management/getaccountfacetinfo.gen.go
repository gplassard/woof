package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAccountFacetInfoCmd = &cobra.Command{
	Use:   "getaccountfacetinfo",
	Short: "Get account facet info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/product-analytics/accounts/facet_info")
		fmt.Println("OperationID: GetAccountFacetInfo")
	},
}

func init() {
	Cmd.AddCommand(GetAccountFacetInfoCmd)
}
