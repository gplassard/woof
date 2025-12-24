package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ConvertJobResultToSignalCmd = &cobra.Command{
	Use:   "convertjobresulttosignal",
	Short: "Convert a job result to a signal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/siem-threat-hunting/jobs/signal_convert")
		fmt.Println("OperationID: ConvertJobResultToSignal")
	},
}

func init() {
	Cmd.AddCommand(ConvertJobResultToSignalCmd)
}
