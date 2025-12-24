package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetMonthlyCostAttributionCmd = &cobra.Command{
	Use:   "getmonthlycostattribution",
	Short: "Get Monthly Cost Attribution",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost_by_tag/monthly_cost_attribution")
		fmt.Println("OperationID: GetMonthlyCostAttribution")
	},
}

func init() {
	Cmd.AddCommand(GetMonthlyCostAttributionCmd)
}
