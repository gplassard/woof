package opsgenie_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteOpsgenieServiceCmd = &cobra.Command{
	Use:   "deleteopsgenieservice",
	Short: "Delete a single service object",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integration/opsgenie/services/{integration_service_id}")
		fmt.Println("OperationID: DeleteOpsgenieService")
	},
}

func init() {
	Cmd.AddCommand(DeleteOpsgenieServiceCmd)
}
