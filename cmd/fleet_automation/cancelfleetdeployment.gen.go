package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CancelFleetDeploymentCmd = &cobra.Command{
	Use:   "cancelfleetdeployment",
	Short: "Cancel a deployment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/unstable/fleet/deployments/{deployment_id}/cancel")
		fmt.Println("OperationID: CancelFleetDeployment")
	},
}

func init() {
	Cmd.AddCommand(CancelFleetDeploymentCmd)
}
