package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListApplicationSecurityWAFCustomRulesCmd = &cobra.Command{
	Use:     "list-application-security-waf-custom-rules",
	Aliases: []string{"list-waf-custom-rules"},
	Short:   "List all WAF custom rules",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.ListApplicationSecurityWAFCustomRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-application-security-waf-custom-rules")

		cmd.Println(cmdutil.FormatJSON(res, "custom_rule"))
	},
}

func init() {
	Cmd.AddCommand(ListApplicationSecurityWAFCustomRulesCmd)
}
