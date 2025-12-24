package data_deletion

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDataDeletionRequestsCmd = &cobra.Command{
	Use:   "getdatadeletionrequests",
	Short: "Gets a list of data deletion requests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/deletion/requests")
		fmt.Println("OperationID: GetDataDeletionRequests")
	},
}

func init() {
	Cmd.AddCommand(GetDataDeletionRequestsCmd)
}
