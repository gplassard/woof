package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCurrentUserApplicationKeysCmd = &cobra.Command{
	Use:   "listcurrentuserapplicationkeys",
	Short: "Get all application keys owned by current user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/current_user/application_keys")
		fmt.Println("OperationID: ListCurrentUserApplicationKeys")
	},
}

func init() {
	Cmd.AddCommand(ListCurrentUserApplicationKeysCmd)
}
