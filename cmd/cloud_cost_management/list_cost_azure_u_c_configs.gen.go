package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCostAzureUCConfigsCmd = &cobra.Command{
	Use:   "list-cost-azure-u-c-configs",
	
	Short: "List Cloud Cost Management Azure configs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCostAzureUCConfigs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-cost-azure-u-c-configs: %v", err)
		}

		cmdutil.PrintJSON(res, "azure_uc_configs")
	},
}

func init() {
	Cmd.AddCommand(ListCostAzureUCConfigsCmd)
}
