package cloud_network_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAggregatedDnsCmd = &cobra.Command{
	Use:   "getaggregateddns",
	Short: "Get all aggregated DNS traffic",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/network/dns/aggregate")
		fmt.Println("OperationID: GetAggregatedDns")
	},
}

func init() {
	Cmd.AddCommand(GetAggregatedDnsCmd)
}
