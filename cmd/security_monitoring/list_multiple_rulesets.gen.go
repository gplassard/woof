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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListMultipleRulesets(client.NewContext(apiKey, appKey, site), datadogV2.GetMultipleRulesetsRequest{})
		cmdutil.HandleError(err, "failed to list-multiple-rulesets")

		cmd.Println(cmdutil.FormatJSON(res, "get_multiple_rulesets_response"))
	},
}

func init() {
	Cmd.AddCommand(ListMultipleRulesetsCmd)
}
