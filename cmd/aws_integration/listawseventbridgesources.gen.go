package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAWSEventBridgeSourcesCmd = &cobra.Command{
	Use:   "listawseventbridgesources",
	Short: "Get all Amazon EventBridge sources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/event_bridge")
		fmt.Println("OperationID: ListAWSEventBridgeSources")
	},
}

func init() {
	Cmd.AddCommand(ListAWSEventBridgeSourcesCmd)
}
