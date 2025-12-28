package application_security

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "create-application-security-waf-custom-rule",
	Aliases: []string{ "create-waf-custom-rule", },
	Short: "Create a WAF custom rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.CreateApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), datadogV2.ApplicationSecurityWafCustomRuleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-application-security-waf-custom-rule")

		cmdutil.PrintJSON(res, "custom_rule")
	},
}

func init() {
	Cmd.AddCommand(CreateApplicationSecurityWafCustomRuleCmd)
}
