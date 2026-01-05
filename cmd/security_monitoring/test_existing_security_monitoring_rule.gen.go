package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var TestExistingSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "test-existing-security-monitoring-rule [rule_id]",
	Aliases: []string{"test-existing-rule"},
	Short:   "Test an existing rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.TestExistingSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringRuleTestRequest{})
		cmdutil.HandleError(err, "failed to test-existing-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(TestExistingSecurityMonitoringRuleCmd)
}
