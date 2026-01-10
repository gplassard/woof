package service_scorecards

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListScorecardOutcomesCmd = &cobra.Command{
	Use:     "list-scorecard-outcomes",
	Aliases: []string{"list-outcomes"},
	Short:   "List all rule outcomes",
	Long: `List all rule outcomes
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#list-scorecard-outcomes`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OutcomesResponse
		var err error

		optionalParams := datadogV2.NewListScorecardOutcomesOptionalParameters()

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

		if cmd.Flags().Changed("fields-outcome") {
			val, _ := cmd.Flags().GetString("fields-outcome")
			optionalParams.WithFieldsOutcome(val)
		}

		if cmd.Flags().Changed("fields-rule") {
			val, _ := cmd.Flags().GetString("fields-rule")
			optionalParams.WithFieldsRule(val)
		}

		if cmd.Flags().Changed("filter-outcome-service-name") {
			val, _ := cmd.Flags().GetString("filter-outcome-service-name")
			optionalParams.WithFilterOutcomeServiceName(val)
		}

		if cmd.Flags().Changed("filter-outcome-state") {
			val, _ := cmd.Flags().GetString("filter-outcome-state")
			optionalParams.WithFilterOutcomeState(val)
		}

		if cmd.Flags().Changed("filter-rule-enabled") {
			val, _ := cmd.Flags().GetString("filter-rule-enabled")
			optionalParams.WithFilterRuleEnabled(val)
		}

		if cmd.Flags().Changed("filter-rule-id") {
			val, _ := cmd.Flags().GetString("filter-rule-id")
			optionalParams.WithFilterRuleId(val)
		}

		if cmd.Flags().Changed("filter-rule-name") {
			val, _ := cmd.Flags().GetString("filter-rule-name")
			optionalParams.WithFilterRuleName(val)
		}

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListScorecardOutcomes(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-scorecard-outcomes")

		fmt.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {

	ListScorecardOutcomesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListScorecardOutcomesCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListScorecardOutcomesCmd.Flags().String("include", "", "Include related rule details in the response.")

	ListScorecardOutcomesCmd.Flags().String("fields-outcome", "", "Return only specified values in the outcome attributes.")

	ListScorecardOutcomesCmd.Flags().String("fields-rule", "", "Return only specified values in the included rule details.")

	ListScorecardOutcomesCmd.Flags().String("filter-outcome-service-name", "", "Filter the outcomes on a specific service name.")

	ListScorecardOutcomesCmd.Flags().String("filter-outcome-state", "", "Filter the outcomes by a specific state.")

	ListScorecardOutcomesCmd.Flags().String("filter-rule-enabled", "", "Filter outcomes on whether a rule is enabled/disabled.")

	ListScorecardOutcomesCmd.Flags().String("filter-rule-id", "", "Filter outcomes based on rule ID.")

	ListScorecardOutcomesCmd.Flags().String("filter-rule-name", "", "Filter outcomes based on rule name.")

	Cmd.AddCommand(ListScorecardOutcomesCmd)
}
