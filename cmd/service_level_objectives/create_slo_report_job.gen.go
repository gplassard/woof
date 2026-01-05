package service_level_objectives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSLOReportJobCmd = &cobra.Command{
	Use:     "create-slo-report-job",
	Aliases: []string{"create-report-job"},
	Short:   "Create a new SLO report",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		res, _, err := api.CreateSLOReportJob(client.NewContext(apiKey, appKey, site), datadogV2.SloReportCreateRequest{})
		cmdutil.HandleError(err, "failed to create-slo-report-job")

		cmd.Println(cmdutil.FormatJSON(res, "service_level_objectives"))
	},
}

func init() {
	Cmd.AddCommand(CreateSLOReportJobCmd)
}
