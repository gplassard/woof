package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:     "create-application-security-waf-custom-rule",
	Aliases: []string{"create-waf-custom-rule"},
	Short:   "Create a WAF custom rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.CreateApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), datadogV2.ApplicationSecurityWafCustomRuleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-application-security-waf-custom-rule")

		cmd.Println(cmdutil.FormatJSON(res, "custom_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateApplicationSecurityWafCustomRuleCmd)
}
