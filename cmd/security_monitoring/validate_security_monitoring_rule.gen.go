package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ValidateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "validate-security-monitoring-rule [payload]",
	Aliases: []string{"validate-rule"},
	Short:   "Validate a detection rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SecurityMonitoringRuleValidatePayload
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.ValidateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-security-monitoring-rule")

	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringRuleCmd)
}
