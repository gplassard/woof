package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CancelThreatHuntingJobCmd = &cobra.Command{
	Use:   "cancelthreathuntingjob",
	Short: "Cancel a threat hunting job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/siem-threat-hunting/jobs/{job_id}/cancel")
		fmt.Println("OperationID: CancelThreatHuntingJob")
	},
}

func init() {
	Cmd.AddCommand(CancelThreatHuntingJobCmd)
}
