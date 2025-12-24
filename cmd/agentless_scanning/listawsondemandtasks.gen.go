package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAwsOnDemandTasksCmd = &cobra.Command{
	Use:   "listawsondemandtasks",
	Short: "List AWS on demand tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/ondemand/aws")
		fmt.Println("OperationID: ListAwsOnDemandTasks")
	},
}

func init() {
	Cmd.AddCommand(ListAwsOnDemandTasksCmd)
}
