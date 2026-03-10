package static_analysis

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCustomRulesetCmd = &cobra.Command{
	Use: "update-custom-ruleset [ruleset_name]",

	Short: "Update Custom Ruleset",
	Long: `Update Custom Ruleset
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#update-custom-ruleset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomRulesetResponse
		var err error

		var body datadogV2.CustomRulesetRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCustomRuleset(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-custom-ruleset")

		fmt.Println(cmdutil.FormatJSON(res, "custom_ruleset"))
	},
}

func init() {

	UpdateCustomRulesetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCustomRulesetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCustomRulesetCmd)
}
