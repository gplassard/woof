package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCostAzureUCConfigsCmd = &cobra.Command{
	Use:   "create_cost_azure_u_c_configs",
	Short: "Create Cloud Cost Management Azure configs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCostAzureUCConfigs(client.NewContext(apiKey, appKey, site), datadogV2.AzureUCConfigPostRequest{})
		if err != nil {
			log.Fatalf("failed to create_cost_azure_u_c_configs: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(CreateCostAzureUCConfigsCmd)
}
