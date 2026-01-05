package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMultipleRulesetsCmd = &cobra.Command{
	Use: "list-multiple-rulesets",

	Short: "Ruleset get multiple",
	Long: `Ruleset get multiple
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-multiple-rulesets`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetMultipleRulesetsResponse
		var err error

		var body datadogV2.GetMultipleRulesetsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.ListMultipleRulesets(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-multiple-rulesets")

		cmd.Println(cmdutil.FormatJSON(res, "get_multiple_rulesets_response"))
	},
}

func init() {

	ListMultipleRulesetsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ListMultipleRulesetsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ListMultipleRulesetsCmd)
}
