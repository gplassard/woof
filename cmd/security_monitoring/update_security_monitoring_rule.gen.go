package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "update-security-monitoring-rule [rule_id]",
	Aliases: []string{"update-rule"},
	Short:   "Update an existing rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringRuleUpdatePayload{})
		cmdutil.HandleError(err, "failed to update-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityMonitoringRuleCmd)
}
