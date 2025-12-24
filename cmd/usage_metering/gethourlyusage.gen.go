package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetHourlyUsageCmd = &cobra.Command{
	Use:   "gethourlyusage",
	Short: "Get hourly usage by product family",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/hourly_usage")
		fmt.Println("OperationID: GetHourlyUsage")
	},
}

func init() {
	Cmd.AddCommand(GetHourlyUsageCmd)
}
