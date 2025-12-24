package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetMonitorUserTemplateCmd = &cobra.Command{
	Use:   "getmonitorusertemplate",
	Short: "Get a monitor user template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/monitor/template/{template_id}")
		fmt.Println("OperationID: GetMonitorUserTemplate")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorUserTemplateCmd)
}
