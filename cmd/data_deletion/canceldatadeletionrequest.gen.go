package data_deletion

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CancelDataDeletionRequestCmd = &cobra.Command{
	Use:   "canceldatadeletionrequest",
	Short: "Cancels a data deletion request",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/deletion/requests/{id}/cancel")
		fmt.Println("OperationID: CancelDataDeletionRequest")
	},
}

func init() {
	Cmd.AddCommand(CancelDataDeletionRequestCmd)
}
