package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringRulesCmd = &cobra.Command{
	Use:     "list-security-monitoring-rules",
	Aliases: []string{"list-rules"},
	Short:   "List rules",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityMonitoringRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-rules")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringRulesCmd)
}
