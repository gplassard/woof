package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RunThreatHuntingJobCmd = &cobra.Command{
	Use:   "runthreathuntingjob",
	Short: "Run a threat hunting job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/siem-threat-hunting/jobs")
		fmt.Println("OperationID: RunThreatHuntingJob")
	},
}

func init() {
	Cmd.AddCommand(RunThreatHuntingJobCmd)
}
