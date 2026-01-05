package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var TestSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "test-security-monitoring-rule",
	Aliases: []string{"test-rule"},
	Short:   "Test a rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.TestSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringRuleTestRequest{})
		cmdutil.HandleError(err, "failed to test-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(TestSecurityMonitoringRuleCmd)
}
