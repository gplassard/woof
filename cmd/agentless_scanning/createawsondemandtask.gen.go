package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAwsOnDemandTaskCmd = &cobra.Command{
	Use:   "createawsondemandtask",
	Short: "Create AWS on demand task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/agentless_scanning/ondemand/aws")
		fmt.Println("OperationID: CreateAwsOnDemandTask")
	},
}

func init() {
	Cmd.AddCommand(CreateAwsOnDemandTaskCmd)
}
