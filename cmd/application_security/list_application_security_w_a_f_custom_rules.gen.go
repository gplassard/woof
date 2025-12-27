package application_security

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListApplicationSecurityWAFCustomRulesCmd = &cobra.Command{
	Use:   "list_application_security_w_a_f_custom_rules",
	Short: "List all WAF custom rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.ListApplicationSecurityWAFCustomRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_application_security_w_a_f_custom_rules: %v", err)
		}

		cmdutil.PrintJSON(res, "application_security")
	},
}

func init() {
	Cmd.AddCommand(ListApplicationSecurityWAFCustomRulesCmd)
}
