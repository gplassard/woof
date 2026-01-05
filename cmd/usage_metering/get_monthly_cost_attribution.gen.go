package usage_metering

import (
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
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetMonthlyCostAttribution(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), args[1])
		cmdutil.HandleError(err, "failed to get-monthly-cost-attribution")

		cmd.Println(cmdutil.FormatJSON(res, "usage_metering"))
	},
}

func init() {
	Cmd.AddCommand(GetMonthlyCostAttributionCmd)
}
