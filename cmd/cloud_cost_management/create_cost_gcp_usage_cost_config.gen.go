package cloud_cost_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "create-cost-gcp-usage-cost-config",
	
	Short: "Create Google Cloud Usage Cost config",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), datadogV2.GCPUsageCostConfigPostRequest{})
		cmdutil.HandleError(err, "failed to create-cost-gcp-usage-cost-config")

		cmdutil.PrintJSON(res, "gcp_uc_config")
	},
}

func init() {
	Cmd.AddCommand(CreateCostGCPUsageCostConfigCmd)
}
