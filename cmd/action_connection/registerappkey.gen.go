package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RegisterAppKeyCmd = &cobra.Command{
	Use:   "registerappkey",
	Short: "Register a new App Key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/actions/app_key_registrations/{app_key_id}")
		fmt.Println("OperationID: RegisterAppKey")
	},
}

func init() {
	Cmd.AddCommand(RegisterAppKeyCmd)
}
