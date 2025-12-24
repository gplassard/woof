package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListFleetDeploymentsCmd = &cobra.Command{
	Use:   "listfleetdeployments",
	Short: "List all deployments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/unstable/fleet/deployments")
		fmt.Println("OperationID: ListFleetDeployments")
	},
}

func init() {
	Cmd.AddCommand(ListFleetDeploymentsCmd)
}
