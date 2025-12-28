package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSecurityMonitoringHistsignalsByJobIdCmd = &cobra.Command{
	Use:   "get_security_monitoring_histsignals_by_job_id [job_id]",
	Short: "Get a job's hist signals",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSecurityMonitoringHistsignalsByJobId(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_security_monitoring_histsignals_by_job_id: %v", err)
		}

		cmdutil.PrintJSON(res, "signal")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringHistsignalsByJobIdCmd)
}
