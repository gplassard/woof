package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateMonitorUserTemplateCmd = &cobra.Command{
	Use:   "createmonitorusertemplate",
	Short: "Create a monitor user template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/monitor/template")
		fmt.Println("OperationID: CreateMonitorUserTemplate")
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorUserTemplateCmd)
}
