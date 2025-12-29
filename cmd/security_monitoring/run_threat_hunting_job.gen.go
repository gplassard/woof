package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RunThreatHuntingJobCmd = &cobra.Command{
	Use:   "run-threat-hunting-job",
	
	Short: "Run a threat hunting job",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.RunThreatHuntingJob(client.NewContext(apiKey, appKey, site), datadogV2.RunThreatHuntingJobRequest{})
		cmdutil.HandleError(err, "failed to run-threat-hunting-job")

		cmd.Println(cmdutil.FormatJSON(res, "historicalDetectionsJob"))
	},
}

func init() {
	Cmd.AddCommand(RunThreatHuntingJobCmd)
}
