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
	Long: `Get a rule's version history
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-rule-version-history`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetRuleVersionHistoryResponse
		var err error

		optionalParams := datadogV2.NewGetRuleVersionHistoryOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetRuleVersionHistory(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-rule-version-history")

		cmd.Println(cmdutil.FormatJSON(res, "GetRuleVersionHistoryResponse"))
	},
}

func init() {

	GetRuleVersionHistoryCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	GetRuleVersionHistoryCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	Cmd.AddCommand(GetRuleVersionHistoryCmd)
}
