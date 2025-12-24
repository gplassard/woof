package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateFleetDeploymentConfigureCmd = &cobra.Command{
	Use:   "createfleetdeploymentconfigure",
	Short: "Create a configuration deployment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/unstable/fleet/deployments/configure")
		fmt.Println("OperationID: CreateFleetDeploymentConfigure")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetDeploymentConfigureCmd)
}
