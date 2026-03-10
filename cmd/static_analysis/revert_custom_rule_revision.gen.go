package static_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RevertCustomRuleRevisionCmd = &cobra.Command{
	Use: "revert-custom-rule-revision [ruleset_name] [rule_name]",

	Short: "Revert Custom Rule Revision",
	Long: `Revert Custom Rule Revision
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#revert-custom-rule-revision`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.RevertCustomRuleRevisionRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.RevertCustomRuleRevision(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to revert-custom-rule-revision")

	},
}

func init() {

	RevertCustomRuleRevisionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	RevertCustomRuleRevisionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(RevertCustomRuleRevisionCmd)
}
