package opsgenie_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOpsgenieServiceCmd = &cobra.Command{
	Use:   "getopsgenieservice",
	Short: "Get a single service object",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/opsgenie/services/{integration_service_id}")
		fmt.Println("OperationID: GetOpsgenieService")
	},
}

func init() {
	Cmd.AddCommand(GetOpsgenieServiceCmd)
}
