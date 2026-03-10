package static_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCustomRulesetCmd = &cobra.Command{
	Use: "delete-custom-ruleset [ruleset_name]",

	Short: "Delete Custom Ruleset",
	Long: `Delete Custom Ruleset
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#delete-custom-ruleset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteCustomRuleset(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-custom-ruleset")

	},
}

func init() {

	Cmd.AddCommand(DeleteCustomRulesetCmd)
}
