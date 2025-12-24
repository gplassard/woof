package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListMonitorUserTemplatesCmd = &cobra.Command{
	Use:   "listmonitorusertemplates",
	Short: "Get all monitor user templates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/monitor/template")
		fmt.Println("OperationID: ListMonitorUserTemplates")
	},
}

func init() {
	Cmd.AddCommand(ListMonitorUserTemplatesCmd)
}
