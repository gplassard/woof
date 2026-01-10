package usage_metering

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"time"

	"github.com/spf13/cobra"
)

var GetHistoricalCostByOrgCmd = &cobra.Command{
	Use: "get-historical-cost-by-org [start_month]",

	Short: "Get historical cost across your account",
	Long: `Get historical cost across your account
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-historical-cost-by-org`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CostByOrgResponse
		var err error

		optionalParams := datadogV2.NewGetHistoricalCostByOrgOptionalParameters()

		if cmd.Flags().Changed("view") {
			val, _ := cmd.Flags().GetString("view")
			optionalParams.WithView(val)
		}

		if cmd.Flags().Changed("end-month") {
			val, _ := cmd.Flags().GetString("end-month")
			optionalParams.WithEndMonth(val)
		}

		if cmd.Flags().Changed("include-connected-accounts") {
			val, _ := cmd.Flags().GetString("include-connected-accounts")
			optionalParams.WithIncludeConnectedAccounts(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetHistoricalCostByOrg(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), *optionalParams)
		cmdutil.HandleError(err, "failed to get-historical-cost-by-org")

		fmt.Println(cmdutil.FormatJSON(res, "cost_by_org"))
	},
}

func init() {

	GetHistoricalCostByOrgCmd.Flags().String("view", "", "String to specify whether cost is broken down at a parent-org level or at the sub-org level. Available views are 'summary' and 'sub-org'.  Defaults to 'summary'.")

	GetHistoricalCostByOrgCmd.Flags().String("end-month", "", "Datetime in ISO-8601 format, UTC, precise to month: '[YYYY-MM]' for cost ending this month.")

	GetHistoricalCostByOrgCmd.Flags().String("include-connected-accounts", "", "Boolean to specify whether to include accounts connected to the current account as partner customers in the Datadog partner network program. Defaults to 'false'. ")

	Cmd.AddCommand(GetHistoricalCostByOrgCmd)
}
