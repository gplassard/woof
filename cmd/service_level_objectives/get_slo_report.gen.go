package service_level_objectives

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSLOReportCmd = &cobra.Command{
	Use:   "get-slo-report [report_id]",
	Aliases: []string{ "get-report", },
	Short: "Get SLO report",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		res, _, err := api.GetSLOReport(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-slo-report")

		cmd.Println(cmdutil.FormatJSON(res, "service_level_objectives"))
	},
}

func init() {
	Cmd.AddCommand(GetSLOReportCmd)
}
