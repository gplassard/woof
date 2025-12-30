package cloud_cost_management

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCostAzureUCConfigsCmd = &cobra.Command{
	Use: "list-cost-azure-uc-configs",

	Short: "List Cloud Cost Management Azure configs",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCostAzureUCConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cost-azure-uc-configs")

		cmd.Println(cmdutil.FormatJSON(res, "azure_uc_configs"))
	},
}

func init() {
	Cmd.AddCommand(ListCostAzureUCConfigsCmd)
}
