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

var GetHourlyUsageCmd = &cobra.Command{
	Use: "get-hourly-usage [filter[timestamp][start]] [filter[product_families]]",

	Short: "Get hourly usage by product family",
	Long: `Get hourly usage by product family
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-hourly-usage`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.HourlyUsageResponse
		var err error

		optionalParams := datadogV2.NewGetHourlyUsageOptionalParameters()

		if cmd.Flags().Changed("filter-timestamp-end") {
			val, _ := cmd.Flags().GetString("filter-timestamp-end")
			optionalParams.WithFilterTimestampEnd(val)
		}

		if cmd.Flags().Changed("filter-include-descendants") {
			val, _ := cmd.Flags().GetString("filter-include-descendants")
			optionalParams.WithFilterIncludeDescendants(val)
		}

		if cmd.Flags().Changed("filter-include-connected-accounts") {
			val, _ := cmd.Flags().GetString("filter-include-connected-accounts")
			optionalParams.WithFilterIncludeConnectedAccounts(val)
		}

		if cmd.Flags().Changed("filter-include-breakdown") {
			val, _ := cmd.Flags().GetString("filter-include-breakdown")
			optionalParams.WithFilterIncludeBreakdown(val)
		}

		if cmd.Flags().Changed("filter-versions") {
			val, _ := cmd.Flags().GetString("filter-versions")
			optionalParams.WithFilterVersions(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("page-next-record-id") {
			val, _ := cmd.Flags().GetString("page-next-record-id")
			optionalParams.WithPageNextRecordId(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetHourlyUsage(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), args[1], *optionalParams)
		cmdutil.HandleError(err, "failed to get-hourly-usage")

		fmt.Println(cmdutil.FormatJSON(res, "usage_timeseries"))
	},
}

func init() {

	GetHourlyUsageCmd.Flags().String("filter-timestamp-end", "", "Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour.")

	GetHourlyUsageCmd.Flags().String("filter-include-descendants", "", "Include child org usage in the response. Defaults to false.")

	GetHourlyUsageCmd.Flags().String("filter-include-connected-accounts", "", "Boolean to specify whether to include accounts connected to the current account as partner customers in the Datadog partner network program. Defaults to false.")

	GetHourlyUsageCmd.Flags().String("filter-include-breakdown", "", "Include breakdown of usage by subcategories where applicable (for product family logs only). Defaults to false.")

	GetHourlyUsageCmd.Flags().String("filter-versions", "", "Comma separated list of product family versions to use in the format 'product_family:version'. For example, 'infra_hosts:1.0.0'. If this parameter is not used, the API will use the latest version of each requested product family. Currently all families have one version '1.0.0'.")

	GetHourlyUsageCmd.Flags().Int64("page-limit", 0, "Maximum number of results to return (between 1 and 500) - defaults to 500 if limit not specified.")

	GetHourlyUsageCmd.Flags().String("page-next-record-id", "", "List following results with a next_record_id provided in the previous query.")

	Cmd.AddCommand(GetHourlyUsageCmd)
}
