package data_deletion

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDataDeletionRequestCmd = &cobra.Command{
	Use:   "createdatadeletionrequest",
	Short: "Creates a data deletion request",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/deletion/data/{product}")
		fmt.Println("OperationID: CreateDataDeletionRequest")
	},
}

func init() {
	Cmd.AddCommand(CreateDataDeletionRequestCmd)
}
