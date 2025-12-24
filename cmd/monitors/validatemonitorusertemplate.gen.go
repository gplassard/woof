package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ValidateMonitorUserTemplateCmd = &cobra.Command{
	Use:   "validatemonitorusertemplate",
	Short: "Validate a monitor user template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/monitor/template/validate")
		fmt.Println("OperationID: ValidateMonitorUserTemplate")
	},
}

func init() {
	Cmd.AddCommand(ValidateMonitorUserTemplateCmd)
}
