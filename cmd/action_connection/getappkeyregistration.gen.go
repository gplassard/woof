package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAppKeyRegistrationCmd = &cobra.Command{
	Use:   "getappkeyregistration",
	Short: "Get an existing App Key Registration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/actions/app_key_registrations/{app_key_id}")
		fmt.Println("OperationID: GetAppKeyRegistration")
	},
}

func init() {
	Cmd.AddCommand(GetAppKeyRegistrationCmd)
}
