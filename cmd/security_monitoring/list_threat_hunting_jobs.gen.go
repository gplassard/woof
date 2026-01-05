package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListThreatHuntingJobsCmd = &cobra.Command{
	Use: "list-threat-hunting-jobs",

	Short: "List threat hunting jobs",
	Long: `List threat hunting jobs
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-threat-hunting-jobs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListThreatHuntingJobsResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.ListThreatHuntingJobs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-threat-hunting-jobs")

		cmd.Println(cmdutil.FormatJSON(res, "historicalDetectionsJob"))
	},
}

func init() {

	Cmd.AddCommand(ListThreatHuntingJobsCmd)
}
