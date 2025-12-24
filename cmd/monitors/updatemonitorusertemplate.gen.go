package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateMonitorUserTemplateCmd = &cobra.Command{
	Use:   "updatemonitorusertemplate",
	Short: "Update a monitor user template to a new version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/monitor/template/{template_id}")
		fmt.Println("OperationID: UpdateMonitorUserTemplate")
	},
}

func init() {
	Cmd.AddCommand(UpdateMonitorUserTemplateCmd)
}
