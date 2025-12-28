package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListThreatHuntingJobsCmd = &cobra.Command{
	Use:   "list_threat_hunting_jobs",
	Short: "List threat hunting jobs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListThreatHuntingJobs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_threat_hunting_jobs: %v", err)
		}

		cmdutil.PrintJSON(res, "historicalDetectionsJob")
	},
}

func init() {
	Cmd.AddCommand(ListThreatHuntingJobsCmd)
}
