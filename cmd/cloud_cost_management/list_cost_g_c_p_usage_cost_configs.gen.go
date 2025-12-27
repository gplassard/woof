package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCostGCPUsageCostConfigsCmd = &cobra.Command{
	Use:   "list_cost_g_c_p_usage_cost_configs",
	Short: "List Google Cloud Usage Cost configs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCostGCPUsageCostConfigs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_cost_g_c_p_usage_cost_configs: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(ListCostGCPUsageCostConfigsCmd)
}
