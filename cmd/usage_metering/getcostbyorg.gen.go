package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCostByOrgCmd = &cobra.Command{
	Use:   "getcostbyorg",
	Short: "Get cost across multi-org account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/cost_by_org")
		fmt.Println("OperationID: GetCostByOrg")
	},
}

func init() {
	Cmd.AddCommand(GetCostByOrgCmd)
}
