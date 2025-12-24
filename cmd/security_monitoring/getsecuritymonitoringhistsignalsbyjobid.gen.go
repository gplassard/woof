package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSecurityMonitoringHistsignalsByJobIdCmd = &cobra.Command{
	Use:   "getsecuritymonitoringhistsignalsbyjobid",
	Short: "Get a job's hist signals",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/siem-threat-hunting/jobs/{job_id}/histsignals")
		fmt.Println("OperationID: GetSecurityMonitoringHistsignalsByJobId")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringHistsignalsByJobIdCmd)
}
