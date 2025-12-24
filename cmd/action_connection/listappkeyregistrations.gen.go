package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAppKeyRegistrationsCmd = &cobra.Command{
	Use:   "listappkeyregistrations",
	Short: "List App Key Registrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/actions/app_key_registrations")
		fmt.Println("OperationID: ListAppKeyRegistrations")
	},
}

func init() {
	Cmd.AddCommand(ListAppKeyRegistrationsCmd)
}
