package service_level_objectives

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateSLOReportJobCmd = &cobra.Command{
	Use:     "create-slo-report-job [payload]",
	Aliases: []string{"create-report-job"},
	Short:   "Create a new SLO report",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SloReportCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewServiceLevelObjectivesApi(client.NewAPIClient())
		res, _, err := api.CreateSLOReportJob(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-slo-report-job")

		cmd.Println(cmdutil.FormatJSON(res, "service_level_objectives"))
	},
}

func init() {
	Cmd.AddCommand(CreateSLOReportJobCmd)
}
