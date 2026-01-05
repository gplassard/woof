package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "update-security-monitoring-rule [rule_id] [payload]",
	Aliases: []string{"update-rule"},
	Short:   "Update an existing rule",
	Long: `Update an existing rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-security-monitoring-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleUpdatePayload
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.UpdateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityMonitoringRuleCmd)
}
