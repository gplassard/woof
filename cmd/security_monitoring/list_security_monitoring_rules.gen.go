package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringRulesCmd = &cobra.Command{
	Use:     "list-security-monitoring-rules",
	Aliases: []string{"list-rules"},
	Short:   "List rules",
	Long: `List rules
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-security-monitoring-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringListRulesResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSecurityMonitoringRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-rules")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	Cmd.AddCommand(ListSecurityMonitoringRulesCmd)
}
