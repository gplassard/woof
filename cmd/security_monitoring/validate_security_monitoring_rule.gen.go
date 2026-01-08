package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "validate-security-monitoring-rule",
	Aliases: []string{"validate-rule"},
	Short:   "Validate a detection rule",
	Long: `Validate a detection rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#validate-security-monitoring-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.SecurityMonitoringRuleValidatePayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.ValidateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-security-monitoring-rule")

	},
}

func init() {

	ValidateSecurityMonitoringRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ValidateSecurityMonitoringRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ValidateSecurityMonitoringRuleCmd)
}
