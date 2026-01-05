package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSuppressionsAffectingRuleCmd = &cobra.Command{
	Use: "get-suppressions-affecting-rule [rule_id]",

	Short: "Get suppressions affecting a specific rule",
	Long: `Get suppressions affecting a specific rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-suppressions-affecting-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSuppressionsResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetSuppressionsAffectingRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-suppressions-affecting-rule")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {

	Cmd.AddCommand(GetSuppressionsAffectingRuleCmd)
}
