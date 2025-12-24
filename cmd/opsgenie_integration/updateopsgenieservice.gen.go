package opsgenie_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateOpsgenieServiceCmd = &cobra.Command{
	Use:   "updateopsgenieservice",
	Short: "Update a single service object",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integration/opsgenie/services/{integration_service_id}")
		fmt.Println("OperationID: UpdateOpsgenieService")
	},
}

func init() {
	Cmd.AddCommand(UpdateOpsgenieServiceCmd)
}
