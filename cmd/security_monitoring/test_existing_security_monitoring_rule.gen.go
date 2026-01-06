package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var TestExistingSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "test-existing-security-monitoring-rule [rule_id]",
	Aliases: []string{"test-existing-rule"},
	Short:   "Test an existing rule",
	Long: `Test an existing rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#test-existing-security-monitoring-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleTestResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleTestRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.TestExistingSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to test-existing-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	TestExistingSecurityMonitoringRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	TestExistingSecurityMonitoringRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(TestExistingSecurityMonitoringRuleCmd)
}
