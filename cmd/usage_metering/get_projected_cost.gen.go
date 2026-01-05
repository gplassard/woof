package usage_metering

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetProjectedCostCmd = &cobra.Command{
	Use: "get-projected-cost",

	Short: "Get projected cost across your account",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetProjectedCost(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-projected-cost")

		cmd.Println(cmdutil.FormatJSON(res, "usage_metering"))
	},
}

func init() {
	Cmd.AddCommand(GetProjectedCostCmd)
}
