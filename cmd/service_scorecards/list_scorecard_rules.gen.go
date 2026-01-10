package service_scorecards

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListScorecardRulesCmd = &cobra.Command{
	Use:     "list-scorecard-rules",
	Aliases: []string{"list-rules"},
	Short:   "List all rules",
	Long: `List all rules
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#list-scorecard-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListRulesResponse
		var err error

		optionalParams := datadogV2.NewListScorecardRulesOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("filter-rule-id") {
			val, _ := cmd.Flags().GetString("filter-rule-id")
			optionalParams.WithFilterRuleId(val)
		}

		if cmd.Flags().Changed("filter-rule-enabled") {
			val, _ := cmd.Flags().GetString("filter-rule-enabled")
			optionalParams.WithFilterRuleEnabled(val)
		}

		if cmd.Flags().Changed("filter-rule-custom") {
			val, _ := cmd.Flags().GetString("filter-rule-custom")
			optionalParams.WithFilterRuleCustom(val)
		}

		if cmd.Flags().Changed("filter-rule-name") {
			val, _ := cmd.Flags().GetString("filter-rule-name")
			optionalParams.WithFilterRuleName(val)
		}

		if cmd.Flags().Changed("filter-rule-description") {
			val, _ := cmd.Flags().GetString("filter-rule-description")
			optionalParams.WithFilterRuleDescription(val)
		}

		if cmd.Flags().Changed("fields-rule") {
			val, _ := cmd.Flags().GetString("fields-rule")
			optionalParams.WithFieldsRule(val)
		}

		if cmd.Flags().Changed("fields-scorecard") {
			val, _ := cmd.Flags().GetString("fields-scorecard")
			optionalParams.WithFieldsScorecard(val)
		}

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListScorecardRules(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-scorecard-rules")

		fmt.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {

	ListScorecardRulesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListScorecardRulesCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListScorecardRulesCmd.Flags().String("include", "", "Include related scorecard details in the response.")

	ListScorecardRulesCmd.Flags().String("filter-rule-id", "", "Filter the rules on a rule ID.")

	ListScorecardRulesCmd.Flags().String("filter-rule-enabled", "", "Filter for enabled rules only.")

	ListScorecardRulesCmd.Flags().String("filter-rule-custom", "", "Filter for custom rules only.")

	ListScorecardRulesCmd.Flags().String("filter-rule-name", "", "Filter rules on the rule name.")

	ListScorecardRulesCmd.Flags().String("filter-rule-description", "", "Filter rules on the rule description.")

	ListScorecardRulesCmd.Flags().String("fields-rule", "", "Return only specific fields in the response for rule attributes.")

	ListScorecardRulesCmd.Flags().String("fields-scorecard", "", "Return only specific fields in the included response for scorecard attributes.")

	Cmd.AddCommand(ListScorecardRulesCmd)
}
