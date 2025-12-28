package application_security

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "get_application_security_waf_custom_rule [custom_rule_id]",
	Short: "Get a WAF custom rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.GetApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_application_security_waf_custom_rule: %v", err)
		}

		cmdutil.PrintJSON(res, "custom_rule")
	},
}

func init() {
	Cmd.AddCommand(GetApplicationSecurityWafCustomRuleCmd)
}
