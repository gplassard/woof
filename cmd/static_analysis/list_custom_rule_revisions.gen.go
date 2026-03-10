package static_analysis

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCustomRuleRevisionsCmd = &cobra.Command{
	Use: "list-custom-rule-revisions [ruleset_name] [rule_name]",

	Short: "List Custom Rule Revisions",
	Long: `List Custom Rule Revisions
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#list-custom-rule-revisions`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomRuleRevisionsResponse
		var err error

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCustomRuleRevisions(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to list-custom-rule-revisions")

		fmt.Println(cmdutil.FormatJSON(res, "custom_rule_revision"))
	},
}

func init() {

	Cmd.AddCommand(ListCustomRuleRevisionsCmd)
}
