package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListFleetAgentVersionsCmd = &cobra.Command{
	Use:   "listfleetagentversions",
	Short: "List all available Agent versions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/unstable/fleet/agent_versions")
		fmt.Println("OperationID: ListFleetAgentVersions")
	},
}

func init() {
	Cmd.AddCommand(ListFleetAgentVersionsCmd)
}
