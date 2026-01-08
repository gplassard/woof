package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCostGCPUsageCostConfigsCmd = &cobra.Command{
	Use: "list-cost-gcp-usage-cost-configs",

	Short: "List Google Cloud Usage Cost configs",
	Long: `List Google Cloud Usage Cost configs
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#list-cost-gcp-usage-cost-configs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPUsageCostConfigsResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCostGCPUsageCostConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cost-gcp-usage-cost-configs")

		fmt.Println(cmdutil.FormatJSON(res, "cost_g_c_p_usage_cost_config"))
	},
}

func init() {

	Cmd.AddCommand(ListCostGCPUsageCostConfigsCmd)
}
