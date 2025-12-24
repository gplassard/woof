package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetThreatHuntingJobCmd = &cobra.Command{
	Use:   "getthreathuntingjob",
	Short: "Get a job's details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/siem-threat-hunting/jobs/{job_id}")
		fmt.Println("OperationID: GetThreatHuntingJob")
	},
}

func init() {
	Cmd.AddCommand(GetThreatHuntingJobCmd)
}
