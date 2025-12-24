package csm_agents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAllCSMServerlessAgentsCmd = &cobra.Command{
	Use:   "listallcsmserverlessagents",
	Short: "Get all CSM Serverless Agents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/csm/onboarding/serverless/agents")
		fmt.Println("OperationID: ListAllCSMServerlessAgents")
	},
}

func init() {
	Cmd.AddCommand(ListAllCSMServerlessAgentsCmd)
}
