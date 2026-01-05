package service_level_objectives

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSLOReportJobStatusCmd = &cobra.Command{
	Use:     "get-slo-report-job-status [report_id]",
	Aliases: []string{"get-report-job-status"},
	Short:   "Get SLO report status",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		res, _, err := api.GetSLOReportJobStatus(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-slo-report-job-status")

		cmd.Println(cmdutil.FormatJSON(res, "service_level_objectives"))
	},
}

func init() {
	Cmd.AddCommand(GetSLOReportJobStatusCmd)
}
