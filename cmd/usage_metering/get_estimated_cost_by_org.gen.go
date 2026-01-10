package usage_metering

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetEstimatedCostByOrgCmd = &cobra.Command{
	Use: "get-estimated-cost-by-org",

	Short: "Get estimated cost across your account",
	Long: `Get estimated cost across your account
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-estimated-cost-by-org`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CostByOrgResponse
		var err error

		optionalParams := datadogV2.NewGetEstimatedCostByOrgOptionalParameters()

		if cmd.Flags().Changed("view") {
			val, _ := cmd.Flags().GetString("view")
			optionalParams.WithView(val)
		}

		if cmd.Flags().Changed("start-month") {
			val, _ := cmd.Flags().GetString("start-month")
			optionalParams.WithStartMonth(val)
		}

		if cmd.Flags().Changed("end-month") {
			val, _ := cmd.Flags().GetString("end-month")
			optionalParams.WithEndMonth(val)
		}

		if cmd.Flags().Changed("start-date") {
			val, _ := cmd.Flags().GetString("start-date")
			optionalParams.WithStartDate(val)
		}

		if cmd.Flags().Changed("end-date") {
			val, _ := cmd.Flags().GetString("end-date")
			optionalParams.WithEndDate(val)
		}

		if cmd.Flags().Changed("include-connected-accounts") {
			val, _ := cmd.Flags().GetString("include-connected-accounts")
			optionalParams.WithIncludeConnectedAccounts(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetEstimatedCostByOrg(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to get-estimated-cost-by-org")

		fmt.Println(cmdutil.FormatJSON(res, "cost_by_org"))
	},
}

func init() {

	GetEstimatedCostByOrgCmd.Flags().String("view", "", "String to specify whether cost is broken down at a parent-org level or at the sub-org level. Available views are 'summary' and 'sub-org'. Defaults to 'summary'.")

	GetEstimatedCostByOrgCmd.Flags().String("start-month", "", "Datetime in ISO-8601 format, UTC, precise to month: '[YYYY-MM]' for cost beginning this month. **Either start_month or start_date should be specified, but not both.** (start_month cannot go beyond two months in the past). Provide an 'end_month' to view month-over-month cost.")

	GetEstimatedCostByOrgCmd.Flags().String("end-month", "", "Datetime in ISO-8601 format, UTC, precise to month: '[YYYY-MM]' for cost ending this month.")

	GetEstimatedCostByOrgCmd.Flags().String("start-date", "", "Datetime in ISO-8601 format, UTC, precise to day: '[YYYY-MM-DD]' for cost beginning this day. **Either start_month or start_date should be specified, but not both.** (start_date cannot go beyond two months in the past). Provide an 'end_date' to view day-over-day cumulative cost.")

	GetEstimatedCostByOrgCmd.Flags().String("end-date", "", "Datetime in ISO-8601 format, UTC, precise to day: '[YYYY-MM-DD]' for cost ending this day.")

	GetEstimatedCostByOrgCmd.Flags().String("include-connected-accounts", "", "Boolean to specify whether to include accounts connected to the current account as partner customers in the Datadog partner network program. Defaults to 'false'. ")

	Cmd.AddCommand(GetEstimatedCostByOrgCmd)
}
