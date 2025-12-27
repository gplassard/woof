package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "delete_security_monitoring_rule [rule_id]",
	Short: "Delete an existing rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DeleteSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_security_monitoring_rule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteSecurityMonitoringRuleCmd)
}
