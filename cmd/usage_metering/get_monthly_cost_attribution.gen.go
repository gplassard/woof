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

var GetMonthlyCostAttributionCmd = &cobra.Command{
	Use: "get-monthly-cost-attribution [start_month] [fields]",

	Short: "Get Monthly Cost Attribution",
	Long: `Get Monthly Cost Attribution
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-monthly-cost-attribution`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonthlyCostAttributionResponse
		var err error

		optionalParams := datadogV2.NewGetMonthlyCostAttributionOptionalParameters()

		if cmd.Flags().Changed("end-month") {
			val, _ := cmd.Flags().GetString("end-month")
			optionalParams.WithEndMonth(val)
		}

		if cmd.Flags().Changed("sort-direction") {
			val, _ := cmd.Flags().GetString("sort-direction")
			optionalParams.WithSortDirection(val)
		}

		if cmd.Flags().Changed("sort-name") {
			val, _ := cmd.Flags().GetString("sort-name")
			optionalParams.WithSortName(val)
		}

		if cmd.Flags().Changed("tag-breakdown-keys") {
			val, _ := cmd.Flags().GetString("tag-breakdown-keys")
			optionalParams.WithTagBreakdownKeys(val)
		}

		if cmd.Flags().Changed("next-record-id") {
			val, _ := cmd.Flags().GetString("next-record-id")
			optionalParams.WithNextRecordId(val)
		}

		if cmd.Flags().Changed("include-descendants") {
			val, _ := cmd.Flags().GetString("include-descendants")
			optionalParams.WithIncludeDescendants(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMonthlyCostAttribution(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), args[1], *optionalParams)
		cmdutil.HandleError(err, "failed to get-monthly-cost-attribution")

		fmt.Println(cmdutil.FormatJSON(res, "cost_by_tag"))
	},
}

func init() {

	GetMonthlyCostAttributionCmd.Flags().String("end-month", "", "Datetime in ISO-8601 format, UTC, precise to month: '[YYYY-MM]' for cost ending this month.")

	GetMonthlyCostAttributionCmd.Flags().String("sort-direction", "", "The direction to sort by: '[desc, asc]'.")

	GetMonthlyCostAttributionCmd.Flags().String("sort-name", "", "The billing dimension to sort by. Always sorted by total cost. Example: 'infra_host'.")

	GetMonthlyCostAttributionCmd.Flags().String("tag-breakdown-keys", "", "Comma separated list of tag keys used to group cost. If no value is provided the cost will not be broken down by tags. To see which tags are available, look for the value of 'tag_config_source' in the API response.")

	GetMonthlyCostAttributionCmd.Flags().String("next-record-id", "", "List following results with a next_record_id provided in the previous query.")

	GetMonthlyCostAttributionCmd.Flags().String("include-descendants", "", "Include child org cost in the response. Defaults to 'true'.")

	Cmd.AddCommand(GetMonthlyCostAttributionCmd)
}
