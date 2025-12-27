package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteThreatHuntingJobCmd = &cobra.Command{
	Use:   "delete_threat_hunting_job [job_id]",
	Short: "Delete an existing job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DeleteThreatHuntingJob(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_threat_hunting_job: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteThreatHuntingJobCmd)
}
