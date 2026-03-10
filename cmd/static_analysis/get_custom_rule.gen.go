package static_analysis

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCustomRuleCmd = &cobra.Command{
	Use: "get-custom-rule [ruleset_name] [rule_name]",

	Short: "Show Custom Rule",
	Long: `Show Custom Rule
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#get-custom-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomRuleResponse
		var err error

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCustomRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-custom-rule")

		fmt.Println(cmdutil.FormatJSON(res, "custom_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetCustomRuleCmd)
}
