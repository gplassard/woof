package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCostAzureUCConfigsCmd = &cobra.Command{
	Use: "list-cost-azure-uc-configs",

	Short: "List Cloud Cost Management Azure configs",
	Long: `List Cloud Cost Management Azure configs
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#list-cost-azure-uc-configs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureUCConfigsResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCostAzureUCConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cost-azure-uc-configs")

		cmd.Println(cmdutil.FormatJSON(res, "azure_uc_configs"))
	},
}

func init() {

	Cmd.AddCommand(ListCostAzureUCConfigsCmd)
}
