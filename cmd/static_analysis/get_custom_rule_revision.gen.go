package static_analysis

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCustomRuleRevisionCmd = &cobra.Command{
	Use: "get-custom-rule-revision [ruleset_name] [rule_name] [id]",

	Short: "Show Custom Rule Revision",
	Long: `Show Custom Rule Revision
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#get-custom-rule-revision`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomRuleRevisionResponse
		var err error

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCustomRuleRevision(client.NewContext(apiKey, appKey, site), args[0], args[1], args[2])
		cmdutil.HandleError(err, "failed to get-custom-rule-revision")

		fmt.Println(cmdutil.FormatJSON(res, "custom_rule_revision"))
	},
}

func init() {

	Cmd.AddCommand(GetCustomRuleRevisionCmd)
}
