package application_security

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "get-application-security-waf-custom-rule [custom_rule_id]",
	Aliases: []string{ "get-waf-custom-rule", },
	Short: "Get a WAF custom rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.GetApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-application-security-waf-custom-rule")

		cmd.Println(cmdutil.FormatJSON(res, "custom_rule"))
	},
}

func init() {
	Cmd.AddCommand(GetApplicationSecurityWafCustomRuleCmd)
}
