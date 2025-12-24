package gcp_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetGCPSTSDelegateCmd = &cobra.Command{
	Use:   "getgcpstsdelegate",
	Short: "List delegate account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/gcp/sts_delegate")
		fmt.Println("OperationID: GetGCPSTSDelegate")
	},
}

func init() {
	Cmd.AddCommand(GetGCPSTSDelegateCmd)
}
