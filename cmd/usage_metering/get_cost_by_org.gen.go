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

var GetCostByOrgCmd = &cobra.Command{
	Use: "get-cost-by-org [start_month]",

	Short: "Get cost across multi-org account",
	Long: `Get cost across multi-org account
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-cost-by-org`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CostByOrgResponse
		var err error

		optionalParams := datadogV2.NewGetCostByOrgOptionalParameters()

		if cmd.Flags().Changed("end-month") {
			val, _ := cmd.Flags().GetString("end-month")
			optionalParams.WithEndMonth(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCostByOrg(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }(), *optionalParams)
		cmdutil.HandleError(err, "failed to get-cost-by-org")

		fmt.Println(cmdutil.FormatJSON(res, "cost_by_org"))
	},
}

func init() {

	GetCostByOrgCmd.Flags().String("end-month", "", "Datetime in ISO-8601 format, UTC, precise to month: '[YYYY-MM]' for cost ending this month.")

	Cmd.AddCommand(GetCostByOrgCmd)
}
