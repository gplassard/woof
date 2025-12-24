package service_level_objectives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSLOReportCmd = &cobra.Command{
	Use:   "getsloreport",
	Short: "Get SLO report",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/slo/report/{report_id}/download")
		fmt.Println("OperationID: GetSLOReport")
	},
}

func init() {
	Cmd.AddCommand(GetSLOReportCmd)
}
