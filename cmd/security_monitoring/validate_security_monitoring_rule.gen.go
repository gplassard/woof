package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ValidateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "validate_security_monitoring_rule",
	Short: "Validate a detection rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.ValidateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringRuleValidatePayload{})
		if err != nil {
			log.Fatalf("failed to validate_security_monitoring_rule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringRuleCmd)
}
