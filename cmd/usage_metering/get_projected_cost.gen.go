package usage_metering

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetProjectedCostCmd = &cobra.Command{
	Use: "get-projected-cost",

	Short: "Get projected cost across your account",
	Long: `Get projected cost across your account
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-projected-cost`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProjectedCostResponse
		var err error

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err = api.GetProjectedCost(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-projected-cost")

		cmd.Println(cmdutil.FormatJSON(res, "projected_cost"))
	},
}

func init() {
	Cmd.AddCommand(GetProjectedCostCmd)
}
