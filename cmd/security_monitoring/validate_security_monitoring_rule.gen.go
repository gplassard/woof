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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.ValidateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringRuleValidatePayload{})
		cmdutil.HandleError(err, "failed to validate-security-monitoring-rule")

	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringRuleCmd)
}
