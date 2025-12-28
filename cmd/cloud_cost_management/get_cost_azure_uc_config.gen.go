package cloud_cost_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var GetCostAzureUCConfigCmd = &cobra.Command{
	Use:   "get-cost-azure-uc-config [cloud_account_id]",
	
	Short: "Get cost Azure UC config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.GetCostAzureUCConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-cost-azure-uc-config")

		cmdutil.PrintJSON(res, "azure_uc_configs")
	},
}

func init() {
	Cmd.AddCommand(GetCostAzureUCConfigCmd)
}
