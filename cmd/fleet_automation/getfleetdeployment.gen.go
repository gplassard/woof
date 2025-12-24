package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetFleetDeploymentCmd = &cobra.Command{
	Use:   "getfleetdeployment",
	Short: "Get a configuration deployment by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/unstable/fleet/deployments/{deployment_id}")
		fmt.Println("OperationID: GetFleetDeployment")
	},
}

func init() {
	Cmd.AddCommand(GetFleetDeploymentCmd)
}
