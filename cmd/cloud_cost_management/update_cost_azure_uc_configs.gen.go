package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"encoding/json"
	"github.com/spf13/cobra"
	"strconv"
)

var UpdateCostAzureUCConfigsCmd = &cobra.Command{
	Use: "update-cost-azure-uc-configs [cloud_account_id] [payload]",

	Short: "Update Cloud Cost Management Azure config",
	Long: `Update Cloud Cost Management Azure config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#update-cost-azure-uc-configs`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureUCConfigPairsResponse
		var err error

		var body datadogV2.AzureUCConfigPatchRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.UpdateCostAzureUCConfigs(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to update-cost-azure-uc-configs")

		cmd.Println(cmdutil.FormatJSON(res, "azure_uc_configs"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCostAzureUCConfigsCmd)
}
