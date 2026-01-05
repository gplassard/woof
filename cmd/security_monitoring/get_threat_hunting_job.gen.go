package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetThreatHuntingJobCmd = &cobra.Command{
	Use: "get-threat-hunting-job [job_id]",

	Short: "Get a job's details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetThreatHuntingJob(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-threat-hunting-job")

		cmd.Println(cmdutil.FormatJSON(res, "historicalDetectionsJob"))
	},
}

func init() {
	Cmd.AddCommand(GetThreatHuntingJobCmd)
}
