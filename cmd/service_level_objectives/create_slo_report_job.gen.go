package service_level_objectives

import (
	"fmt"
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
	Long: `Create a new SLO report
Documentation: https://docs.datadoghq.com/api/latest/service-level-objectives/#create-slo-report-job`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SLOReportPostResponse
		var err error

		var body datadogV2.SloReportCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSLOReportJob(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-slo-report-job")

		fmt.Println(cmdutil.FormatJSON(res, "s_l_o_report_job"))
	},
}

func init() {

	CreateSLOReportJobCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSLOReportJobCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSLOReportJobCmd)
}
