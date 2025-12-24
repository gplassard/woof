package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteApplicationKeyCmd = &cobra.Command{
	Use:   "deleteapplicationkey",
	Short: "Delete an application key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/application_keys/{app_key_id}")
		fmt.Println("OperationID: DeleteApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(DeleteApplicationKeyCmd)
}
