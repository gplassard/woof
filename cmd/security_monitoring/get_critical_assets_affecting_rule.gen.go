package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCriticalAssetsAffectingRuleCmd = &cobra.Command{
	Use: "get-critical-assets-affecting-rule [rule_id]",

	Short: "Get critical assets affecting a specific rule",
	Long: `Get critical assets affecting a specific rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-critical-assets-affecting-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringCriticalAssetsResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCriticalAssetsAffectingRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-critical-assets-affecting-rule")

		fmt.Println(cmdutil.FormatJSON(res, "critical_assets"))
	},
}

func init() {

	Cmd.AddCommand(GetCriticalAssetsAffectingRuleCmd)
}
