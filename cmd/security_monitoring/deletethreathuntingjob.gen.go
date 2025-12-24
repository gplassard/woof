package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteThreatHuntingJobCmd = &cobra.Command{
	Use:   "deletethreathuntingjob",
	Short: "Delete an existing job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/siem-threat-hunting/jobs/{job_id}")
		fmt.Println("OperationID: DeleteThreatHuntingJob")
	},
}

func init() {
	Cmd.AddCommand(DeleteThreatHuntingJobCmd)
}
