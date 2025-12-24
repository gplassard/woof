package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateFleetDeploymentUpgradeCmd = &cobra.Command{
	Use:   "createfleetdeploymentupgrade",
	Short: "Upgrade hosts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/unstable/fleet/deployments/upgrade")
		fmt.Println("OperationID: CreateFleetDeploymentUpgrade")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetDeploymentUpgradeCmd)
}
