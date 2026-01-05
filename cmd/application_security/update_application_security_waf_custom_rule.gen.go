package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:     "update-application-security-waf-custom-rule [custom_rule_id]",
	Aliases: []string{"update-waf-custom-rule"},
	Short:   "Update a WAF Custom Rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.UpdateApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationSecurityWafCustomRuleUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-application-security-waf-custom-rule")

		cmd.Println(cmdutil.FormatJSON(res, "custom_rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationSecurityWafCustomRuleCmd)
}
