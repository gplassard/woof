package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DetachCaseCmd = &cobra.Command{
	Use:   "detach_case",
	Short: "Detach security findings from their case",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DetachCase(client.NewContext(apiKey, appKey, site), datadogV2.DetachCaseRequest{})
		if err != nil {
			log.Fatalf("failed to detach_case: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DetachCaseCmd)
}
