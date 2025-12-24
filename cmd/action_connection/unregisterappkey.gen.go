package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UnregisterAppKeyCmd = &cobra.Command{
	Use:   "unregisterappkey",
	Short: "Unregister an App Key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/actions/app_key_registrations/{app_key_id}")
		fmt.Println("OperationID: UnregisterAppKey")
	},
}

func init() {
	Cmd.AddCommand(UnregisterAppKeyCmd)
}
