package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "delete-security-monitoring-rule [rule_id]",
	Aliases: []string{ "delete-rule", },
	Short: "Delete an existing rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DeleteSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-security-monitoring-rule")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteSecurityMonitoringRuleCmd)
}
