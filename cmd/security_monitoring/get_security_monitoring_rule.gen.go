package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "get-security-monitoring-rule [rule_id]",
	Aliases: []string{"get-rule"},
	Short:   "Get a rule's details",
	Long: `Get a rule's details
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-security-monitoring-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringRuleResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-security-monitoring-rule")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetSecurityMonitoringRuleCmd)
}
