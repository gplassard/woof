package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:     "create-security-monitoring-rule",
	Aliases: []string{"create-rule"},
	Short:   "Create a detection rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateSecurityMonitoringRule(client.NewContext(apiKey, appKey, site), datadogV2.SecurityMonitoringRuleCreatePayload{})
		cmdutil.HandleError(err, "failed to create-security-monitoring-rule")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityMonitoringRuleCmd)
}
