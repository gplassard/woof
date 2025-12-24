package csm_agents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAllCSMAgentsCmd = &cobra.Command{
	Use:   "listallcsmagents",
	Short: "Get all CSM Agents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/csm/onboarding/agents")
		fmt.Println("OperationID: ListAllCSMAgents")
	},
}

func init() {
	Cmd.AddCommand(ListAllCSMAgentsCmd)
}
