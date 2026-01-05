package usage_metering

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetEstimatedCostByOrgCmd = &cobra.Command{
	Use: "get-estimated-cost-by-org",

	Short: "Get estimated cost across your account",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetEstimatedCostByOrg(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-estimated-cost-by-org")

		cmd.Println(cmdutil.FormatJSON(res, "usage_metering"))
	},
}

func init() {
	Cmd.AddCommand(GetEstimatedCostByOrgCmd)
}
