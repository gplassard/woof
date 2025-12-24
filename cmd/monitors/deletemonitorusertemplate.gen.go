package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteMonitorUserTemplateCmd = &cobra.Command{
	Use:   "deletemonitorusertemplate",
	Short: "Delete a monitor user template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/monitor/template/{template_id}")
		fmt.Println("OperationID: DeleteMonitorUserTemplate")
	},
}

func init() {
	Cmd.AddCommand(DeleteMonitorUserTemplateCmd)
}
