package service_level_objectives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSLOReportJobCmd = &cobra.Command{
	Use:   "createsloreportjob",
	Short: "Create a new SLO report",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/slo/report")
		fmt.Println("OperationID: CreateSLOReportJob")
	},
}

func init() {
	Cmd.AddCommand(CreateSLOReportJobCmd)
}
