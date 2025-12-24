package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAWSEventBridgeSourceCmd = &cobra.Command{
	Use:   "createawseventbridgesource",
	Short: "Create an Amazon EventBridge source",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/aws/event_bridge")
		fmt.Println("OperationID: CreateAWSEventBridgeSource")
	},
}

func init() {
	Cmd.AddCommand(CreateAWSEventBridgeSourceCmd)
}
