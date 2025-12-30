package cloud_cost_management

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var GetCostGCPUsageCostConfigCmd = &cobra.Command{
	Use: "get-cost-gcp-usage-cost-config [cloud_account_id]",

	Short: "Get Google Cloud Usage Cost config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.GetCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-cost-gcp-usage-cost-config")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_uc_config"))
	},
}

func init() {
	Cmd.AddCommand(GetCostGCPUsageCostConfigCmd)
}
