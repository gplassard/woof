package service_level_objectives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSLOReportJobStatusCmd = &cobra.Command{
	Use:   "getsloreportjobstatus",
	Short: "Get SLO report status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/slo/report/{report_id}/status")
		fmt.Println("OperationID: GetSLOReportJobStatus")
	},
}

func init() {
	Cmd.AddCommand(GetSLOReportJobStatusCmd)
}
