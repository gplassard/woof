package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "createcurrentuserapplicationkey",
	Short: "Create an application key for current user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/current_user/application_keys")
		fmt.Println("OperationID: CreateCurrentUserApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(CreateCurrentUserApplicationKeyCmd)
}
