package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFindingsCmd = &cobra.Command{
	Use: "list-findings",

	Short: "List findings",
	Long: `List findings
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-findings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListFindingsResponse
		var err error

		optionalParams := datadogV2.NewListFindingsOptionalParameters()

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("snapshot-timestamp") {
			val, _ := cmd.Flags().GetInt64("snapshot-timestamp")
			optionalParams.WithSnapshotTimestamp(val)
		}

		if cmd.Flags().Changed("page-cursor") {
			val, _ := cmd.Flags().GetString("page-cursor")
			optionalParams.WithPageCursor(val)
		}

		if cmd.Flags().Changed("filter-tags") {
			val, _ := cmd.Flags().GetString("filter-tags")
			optionalParams.WithFilterTags(val)
		}

		if cmd.Flags().Changed("filter-evaluation-changed-at") {
			val, _ := cmd.Flags().GetString("filter-evaluation-changed-at")
			optionalParams.WithFilterEvaluationChangedAt(val)
		}

		if cmd.Flags().Changed("filter-muted") {
			val, _ := cmd.Flags().GetString("filter-muted")
			optionalParams.WithFilterMuted(val)
		}

		if cmd.Flags().Changed("filter-rule-id") {
			val, _ := cmd.Flags().GetString("filter-rule-id")
			optionalParams.WithFilterRuleId(val)
		}

		if cmd.Flags().Changed("filter-rule-name") {
			val, _ := cmd.Flags().GetString("filter-rule-name")
			optionalParams.WithFilterRuleName(val)
		}

		if cmd.Flags().Changed("filter-resource-type") {
			val, _ := cmd.Flags().GetString("filter-resource-type")
			optionalParams.WithFilterResourceType(val)
		}

		if cmd.Flags().Changed("filter-@resource-id") {
			val, _ := cmd.Flags().GetString("filter-@resource-id")
			optionalParams.WithFilterResourceId(val)
		}

		if cmd.Flags().Changed("filter-discovery-timestamp") {
			val, _ := cmd.Flags().GetString("filter-discovery-timestamp")
			optionalParams.WithFilterDiscoveryTimestamp(val)
		}

		if cmd.Flags().Changed("filter-evaluation") {
			val, _ := cmd.Flags().GetString("filter-evaluation")
			optionalParams.WithFilterEvaluation(val)
		}

		if cmd.Flags().Changed("filter-status") {
			val, _ := cmd.Flags().GetString("filter-status")
			optionalParams.WithFilterStatus(val)
		}

		if cmd.Flags().Changed("filter-vulnerability-type") {
			val, _ := cmd.Flags().GetString("filter-vulnerability-type")
			optionalParams.WithFilterVulnerabilityType(val)
		}

		if cmd.Flags().Changed("detailed-findings") {
			val, _ := cmd.Flags().GetString("detailed-findings")
			optionalParams.WithDetailedFindings(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFindings(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-findings")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	ListFindingsCmd.Flags().Int64("page-limit", 0, "Limit the number of findings returned. Must be <= 1000.")

	ListFindingsCmd.Flags().Int64("snapshot-timestamp", 0, "Return findings for a given snapshot of time (Unix ms).")

	ListFindingsCmd.Flags().String("page-cursor", "", "Return the next page of findings pointed to by the cursor.")

	ListFindingsCmd.Flags().String("filter-tags", "", "Return findings that have these associated tags (repeatable).")

	ListFindingsCmd.Flags().String("filter-evaluation-changed-at", "", "Return findings that have changed from pass to fail or vice versa on a specified date (Unix ms) or date range (using comparison operators).")

	ListFindingsCmd.Flags().String("filter-muted", "", "Set to 'true' to return findings that are muted. Set to 'false' to return unmuted findings.")

	ListFindingsCmd.Flags().String("filter-rule-id", "", "Return findings for the specified rule ID.")

	ListFindingsCmd.Flags().String("filter-rule-name", "", "Return findings for the specified rule.")

	ListFindingsCmd.Flags().String("filter-resource-type", "", "Return only findings for the specified resource type.")

	ListFindingsCmd.Flags().String("filter-@resource-id", "", "Return only findings for the specified resource id.")

	ListFindingsCmd.Flags().String("filter-discovery-timestamp", "", "Return findings that were found on a specified date (Unix ms) or date range (using comparison operators).")

	ListFindingsCmd.Flags().String("filter-evaluation", "", "Return only 'pass' or 'fail' findings.")

	ListFindingsCmd.Flags().String("filter-status", "", "Return only findings with the specified status.")

	ListFindingsCmd.Flags().String("filter-vulnerability-type", "", "Return findings that match the selected vulnerability types (repeatable).")

	ListFindingsCmd.Flags().String("detailed-findings", "", "Return additional fields for some findings.")

	Cmd.AddCommand(ListFindingsCmd)
}
