package cloud_cost_management

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCostAzureUCConfigsCmd = &cobra.Command{
	Use: "create-cost-azure-uc-configs",

	Short: "Create Cloud Cost Management Azure configs",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCostAzureUCConfigs(client.NewContext(apiKey, appKey, site), datadogV2.AzureUCConfigPostRequest{})
		cmdutil.HandleError(err, "failed to create-cost-azure-uc-configs")

		cmd.Println(cmdutil.FormatJSON(res, "azure_uc_configs"))
	},
}

func init() {
	Cmd.AddCommand(CreateCostAzureUCConfigsCmd)
}
