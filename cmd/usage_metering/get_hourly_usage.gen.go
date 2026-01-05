package usage_metering

import (
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

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err = api.GetHourlyUsage(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), args[1])
		cmdutil.HandleError(err, "failed to get-hourly-usage")

		cmd.Println(cmdutil.FormatJSON(res, "usage_timeseries"))
	},
}

func init() {

	Cmd.AddCommand(GetHourlyUsageCmd)
}
