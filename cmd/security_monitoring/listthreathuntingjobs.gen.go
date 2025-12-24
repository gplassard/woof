package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListThreatHuntingJobsCmd = &cobra.Command{
	Use:   "listthreathuntingjobs",
	Short: "List threat hunting jobs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/siem-threat-hunting/jobs")
		fmt.Println("OperationID: ListThreatHuntingJobs")
	},
}

func init() {
	Cmd.AddCommand(ListThreatHuntingJobsCmd)
}
