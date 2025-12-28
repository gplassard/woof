package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var DeleteCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "delete-cost-g-c-p-usage-cost-config [cloud_account_id]",
	
	Short: "Delete Google Cloud Usage Cost config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.DeleteCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		if err != nil {
			log.Fatalf("failed to delete-cost-g-c-p-usage-cost-config: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCostGCPUsageCostConfigCmd)
}
