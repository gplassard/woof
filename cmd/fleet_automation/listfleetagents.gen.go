package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListFleetAgentsCmd = &cobra.Command{
	Use:   "listfleetagents",
	Short: "List all Datadog Agents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/unstable/fleet/agents")
		fmt.Println("OperationID: ListFleetAgents")
	},
}

func init() {
	Cmd.AddCommand(ListFleetAgentsCmd)
}
