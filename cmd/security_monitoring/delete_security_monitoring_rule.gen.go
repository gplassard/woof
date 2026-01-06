package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "delete-security-monitoring-rule [rule_id]",
	Aliases: []string{"delete-rule"},
	Short:   "Delete an existing rule",
	Long: `Delete an existing rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#delete-security-monitoring-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-security-monitoring-rule")

	},
}

func init() {

	Cmd.AddCommand(DeleteSecurityMonitoringRuleCmd)
}
