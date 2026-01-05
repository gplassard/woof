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
	Long: `Test a rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#test-security-monitoring-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleTestResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleTestRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.TestSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to test-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	TestSecurityMonitoringRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	TestSecurityMonitoringRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(TestSecurityMonitoringRuleCmd)
}
