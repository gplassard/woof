package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRuleVersionHistoryCmd = &cobra.Command{
	Use: "get-rule-version-history [rule_id]",

	Short: "Get a rule's version history",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetRuleVersionHistoryResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetRuleVersionHistory(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-rule-version-history")

		cmd.Println(cmdutil.FormatJSON(res, "GetRuleVersionHistoryResponse"))
	},
}

func init() {
	Cmd.AddCommand(GetRuleVersionHistoryCmd)
}
