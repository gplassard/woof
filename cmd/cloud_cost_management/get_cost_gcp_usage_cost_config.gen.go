package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var GetCostGCPUsageCostConfigCmd = &cobra.Command{
	Use: "get-cost-gcp-usage-cost-config [cloud_account_id]",

	Short: "Get Google Cloud Usage Cost config",
	Long: `Get Google Cloud Usage Cost config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#get-cost-gcp-usage-cost-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GcpUcConfigResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-cost-gcp-usage-cost-config")

		fmt.Println(cmdutil.FormatJSON(res, "cost_g_c_p_usage_cost_config"))
	},
}

func init() {

	Cmd.AddCommand(GetCostGCPUsageCostConfigCmd)
}
