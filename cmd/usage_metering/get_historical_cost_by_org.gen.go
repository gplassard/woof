package usage_metering

import (
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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetHistoricalCostByOrg(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }())
		cmdutil.HandleError(err, "failed to get-historical-cost-by-org")

		cmd.Println(cmdutil.FormatJSON(res, "usage_metering"))
	},
}

func init() {
	Cmd.AddCommand(GetHistoricalCostByOrgCmd)
}
