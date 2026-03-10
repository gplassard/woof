package static_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCustomRuleCmd = &cobra.Command{
	Use: "delete-custom-rule [ruleset_name] [rule_name]",

	Short: "Delete Custom Rule",
	Long: `Delete Custom Rule
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#delete-custom-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteCustomRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-custom-rule")

	},
}

func init() {

	Cmd.AddCommand(DeleteCustomRuleCmd)
}
