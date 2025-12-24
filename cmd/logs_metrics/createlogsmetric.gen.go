package logs_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateLogsMetricCmd = &cobra.Command{
	Use:   "createlogsmetric",
	Short: "Create a log-based metric",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateLogsMetric(client.NewContext(apiKey, appKey, site), datadogV2.LogsMetricCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createlogsmetric: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_metrics")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsMetricCmd)
}
