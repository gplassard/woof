package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetFleetAgentInfoCmd = &cobra.Command{
	Use:   "getfleetagentinfo",
	Short: "Get detailed information about an agent",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/unstable/fleet/agents/{agent_key}")
		fmt.Println("OperationID: GetFleetAgentInfo")
	},
}

func init() {
	Cmd.AddCommand(GetFleetAgentInfoCmd)
}
