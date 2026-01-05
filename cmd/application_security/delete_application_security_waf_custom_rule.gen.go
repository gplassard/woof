package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:     "delete-application-security-waf-custom-rule [custom_rule_id]",
	Aliases: []string{"delete-waf-custom-rule"},
	Short:   "Delete a WAF Custom Rule",
	Long: `Delete a WAF Custom Rule
Documentation: https://docs.datadoghq.com/api/latest/application-security/#delete-application-security-waf-custom-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		_, err = api.DeleteApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-application-security-waf-custom-rule")

	},
}

func init() {
	Cmd.AddCommand(DeleteApplicationSecurityWafCustomRuleCmd)
}
