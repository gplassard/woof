package security_monitoring

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteThreatHuntingJobCmd = &cobra.Command{
	Use: "delete-threat-hunting-job [job_id]",

	Short: "Delete an existing job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DeleteThreatHuntingJob(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-threat-hunting-job")

	},
}

func init() {
	Cmd.AddCommand(DeleteThreatHuntingJobCmd)
}
