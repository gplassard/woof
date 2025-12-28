package cloud_cost_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var UpdateCostAzureUCConfigsCmd = &cobra.Command{
	Use:   "update-cost-azure-uc-configs [cloud_account_id]",
	
	Short: "Update Cloud Cost Management Azure config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCostAzureUCConfigs(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.AzureUCConfigPatchRequest{})
		cmdutil.HandleError(err, "failed to update-cost-azure-uc-configs")

		cmdutil.PrintJSON(res, "azure_uc_configs")
	},
}

func init() {
	Cmd.AddCommand(UpdateCostAzureUCConfigsCmd)
}
