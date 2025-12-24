package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var EditSecurityMonitoringSignalAssigneeCmd = &cobra.Command{
	Use:   "editsecuritymonitoringsignalassignee",
	Short: "Modify the triage assignee of a security signal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security_monitoring/signals/{signal_id}/assignee")
		fmt.Println("OperationID: EditSecurityMonitoringSignalAssignee")
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalAssigneeCmd)
}
