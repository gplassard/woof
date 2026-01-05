package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ListMultipleRulesetsCmd = &cobra.Command{
	Use: "list-multiple-rulesets [payload]",

	Short: "Ruleset get multiple",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetMultipleRulesetsResponse
		var err error

		var body datadogV2.GetMultipleRulesetsRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.ListMultipleRulesets(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-multiple-rulesets")

		cmd.Println(cmdutil.FormatJSON(res, "get_multiple_rulesets_response"))
	},
}

func init() {
	Cmd.AddCommand(ListMultipleRulesetsCmd)
}
