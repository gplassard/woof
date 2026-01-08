package application_security

import (
	"fmt"
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
	Long: `Update a WAF Custom Rule
Documentation: https://docs.datadoghq.com/api/latest/application-security/#update-application-security-waf-custom-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationSecurityWafCustomRuleResponse
		var err error

		var body datadogV2.ApplicationSecurityWafCustomRuleUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-application-security-waf-custom-rule")

		fmt.Println(cmdutil.FormatJSON(res, "application_security_waf_custom_rule"))
	},
}

func init() {

	UpdateApplicationSecurityWafCustomRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateApplicationSecurityWafCustomRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateApplicationSecurityWafCustomRuleCmd)
}
