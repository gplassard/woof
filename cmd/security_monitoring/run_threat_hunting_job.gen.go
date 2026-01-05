package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var RunThreatHuntingJobCmd = &cobra.Command{
	Use: "run-threat-hunting-job [payload]",

	Short: "Run a threat hunting job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.RunThreatHuntingJobRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.RunThreatHuntingJob(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to run-threat-hunting-job")

		cmd.Println(cmdutil.FormatJSON(res, "historicalDetectionsJob"))
	},
}

func init() {
	Cmd.AddCommand(RunThreatHuntingJobCmd)
}
