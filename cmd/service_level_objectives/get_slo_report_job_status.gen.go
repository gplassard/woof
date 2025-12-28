package service_level_objectives

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSLOReportJobStatusCmd = &cobra.Command{
	Use:   "get-slo-report-job-status [report_id]",
	
	Short: "Get SLO report status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		res, _, err := api.GetSLOReportJobStatus(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-slo-report-job-status: %v", err)
		}

		cmdutil.PrintJSON(res, "service_level_objectives")
	},
}

func init() {
	Cmd.AddCommand(GetSLOReportJobStatusCmd)
}
