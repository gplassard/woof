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

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetEstimatedCostByOrg(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-estimated-cost-by-org")

		fmt.Println(cmdutil.FormatJSON(res, "cost_by_org"))
	},
}

func init() {

	Cmd.AddCommand(GetEstimatedCostByOrgCmd)
}
