package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ValidateExistingMonitorUserTemplateCmd = &cobra.Command{
	Use:   "validateexistingmonitorusertemplate",
	Short: "Validate an existing monitor user template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/monitor/template/{template_id}/validate")
		fmt.Println("OperationID: ValidateExistingMonitorUserTemplate")
	},
}

func init() {
	Cmd.AddCommand(ValidateExistingMonitorUserTemplateCmd)
}
