package static_analysis

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCustomRuleCmd = &cobra.Command{
	Use: "create-custom-rule [ruleset_name]",

	Short: "Create Custom Rule",
	Long: `Create Custom Rule
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#create-custom-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomRuleResponse
		var err error

		var body datadogV2.CustomRuleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCustomRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-custom-rule")

		fmt.Println(cmdutil.FormatJSON(res, "custom_rule"))
	},
}

func init() {

	CreateCustomRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCustomRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCustomRuleCmd)
}
