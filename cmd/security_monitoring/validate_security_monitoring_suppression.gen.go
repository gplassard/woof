package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ValidateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "validate-security-monitoring-suppression",
	Aliases: []string{ "validate-suppression", },
	Short: "Validate a suppression rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.ValidateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringSuppressionCreateRequest{})
		if err != nil {
			log.Fatalf("failed to validate-security-monitoring-suppression: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringSuppressionCmd)
}
