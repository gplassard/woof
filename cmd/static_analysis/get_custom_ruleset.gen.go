package static_analysis

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCustomRulesetCmd = &cobra.Command{
	Use: "get-custom-ruleset [ruleset_name]",

	Short: "Show Custom Ruleset",
	Long: `Show Custom Ruleset
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#get-custom-ruleset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomRulesetResponse
		var err error

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCustomRuleset(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-custom-ruleset")

		fmt.Println(cmdutil.FormatJSON(res, "custom_ruleset"))
	},
}

func init() {

	Cmd.AddCommand(GetCustomRulesetCmd)
}
