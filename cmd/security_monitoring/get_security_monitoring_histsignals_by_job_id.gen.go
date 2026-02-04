package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecurityMonitoringHistsignalsByJobIdCmd = &cobra.Command{
	Use:     "get-security-monitoring-histsignals-by-job-id [job_id]",
	Aliases: []string{"get-histsignals-by-job-id"},
	Short:   "Get a job's hist signals",
	Long: `Get a job's hist signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-security-monitoring-histsignals-by-job-id`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSecurityMonitoringHistsignalsByJobId(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-security-monitoring-histsignals-by-job-id")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring_histsignals_by_job_id"))
	},
}

func init() {

	Cmd.AddCommand(GetSecurityMonitoringHistsignalsByJobIdCmd)
}
