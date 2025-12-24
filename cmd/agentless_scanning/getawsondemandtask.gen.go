package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAwsOnDemandTaskCmd = &cobra.Command{
	Use:   "getawsondemandtask",
	Short: "Get AWS on demand task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/ondemand/aws/{task_id}")
		fmt.Println("OperationID: GetAwsOnDemandTask")
	},
}

func init() {
	Cmd.AddCommand(GetAwsOnDemandTaskCmd)
}
