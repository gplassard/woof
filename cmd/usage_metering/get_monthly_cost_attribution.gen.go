package usage_metering

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	"time"
	
	"github.com/spf13/cobra"
	
)

var GetMonthlyCostAttributionCmd = &cobra.Command{
	Use:   "get-monthly-cost-attribution [start_month] [fields]",
	
	Short: "Get Monthly Cost Attribution",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetMonthlyCostAttribution(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), args[1])
		if err != nil {
			log.Fatalf("failed to get-monthly-cost-attribution: %v", err)
		}

		cmdutil.PrintJSON(res, "usage_metering")
	},
}

func init() {
	Cmd.AddCommand(GetMonthlyCostAttributionCmd)
}
