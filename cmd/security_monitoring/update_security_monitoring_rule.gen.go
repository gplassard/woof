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
	Long: `Update an existing rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-security-monitoring-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleResponse
		var err error

		var body datadogV2.SecurityMonitoringRuleUpdatePayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.UpdateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	UpdateSecurityMonitoringRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateSecurityMonitoringRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateSecurityMonitoringRuleCmd)
}
