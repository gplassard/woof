package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateApplicationKeyCmd = &cobra.Command{
	Use:   "updateapplicationkey",
	Short: "Edit an application key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/application_keys/{app_key_id}")
		fmt.Println("OperationID: UpdateApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationKeyCmd)
}
