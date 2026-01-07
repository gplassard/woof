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

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMonthlyCostAttribution(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), args[1])
		cmdutil.HandleError(err, "failed to get-monthly-cost-attribution")

		fmt.Println(cmdutil.FormatJSON(res, "cost_by_tag"))
	},
}

func init() {

	Cmd.AddCommand(GetMonthlyCostAttributionCmd)
}
