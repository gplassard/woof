package cloud_cost_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCostGCPUsageCostConfigsCmd = &cobra.Command{
	Use: "list-cost-gcp-usage-cost-configs",

	Short: "List Google Cloud Usage Cost configs",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCostGCPUsageCostConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cost-gcp-usage-cost-configs")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_uc_config"))
	},
}

func init() {
	Cmd.AddCommand(ListCostGCPUsageCostConfigsCmd)
}
