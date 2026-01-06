package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RunThreatHuntingJobCmd = &cobra.Command{
	Use: "run-threat-hunting-job",

	Short: "Run a threat hunting job",
	Long: `Run a threat hunting job
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#run-threat-hunting-job`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.JobCreateResponse
		var err error

		var body datadogV2.RunThreatHuntingJobRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.RunThreatHuntingJob(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to run-threat-hunting-job")

		cmd.Println(cmdutil.FormatJSON(res, "historicalDetectionsJob"))
	},
}

func init() {

	RunThreatHuntingJobCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	RunThreatHuntingJobCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(RunThreatHuntingJobCmd)
}
