package application_security

import (
	"fmt"
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
	Long: `Create a WAF custom rule
Documentation: https://docs.datadoghq.com/api/latest/application-security/#create-application-security-waf-custom-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationSecurityWafCustomRuleResponse
		var err error

		var body datadogV2.ApplicationSecurityWafCustomRuleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateApplicationSecurityWafCustomRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-application-security-waf-custom-rule")

		fmt.Println(cmdutil.FormatJSON(res, "custom_rule"))
	},
}

func init() {

	CreateApplicationSecurityWafCustomRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateApplicationSecurityWafCustomRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateApplicationSecurityWafCustomRuleCmd)
}
