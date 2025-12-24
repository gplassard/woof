package cloud_network_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAggregatedConnectionsCmd = &cobra.Command{
	Use:   "getaggregatedconnections",
	Short: "Get all aggregated connections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/network/connections/aggregate")
		fmt.Println("OperationID: GetAggregatedConnections")
	},
}

func init() {
	Cmd.AddCommand(GetAggregatedConnectionsCmd)
}
