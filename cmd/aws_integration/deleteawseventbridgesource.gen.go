package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAWSEventBridgeSourceCmd = &cobra.Command{
	Use:   "deleteawseventbridgesource",
	Short: "Delete an Amazon EventBridge source",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integration/aws/event_bridge")
		fmt.Println("OperationID: DeleteAWSEventBridgeSource")
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSEventBridgeSourceCmd)
}
