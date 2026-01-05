package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CancelThreatHuntingJobCmd = &cobra.Command{
	Use: "cancel-threat-hunting-job [job_id]",

	Short: "Cancel a threat hunting job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err = api.CancelThreatHuntingJob(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to cancel-threat-hunting-job")

	},
}

func init() {
	Cmd.AddCommand(CancelThreatHuntingJobCmd)
}
