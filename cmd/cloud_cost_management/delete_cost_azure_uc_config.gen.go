package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var DeleteCostAzureUCConfigCmd = &cobra.Command{
	Use: "delete-cost-azure-uc-config [cloud_account_id]",

	Short: "Delete Cloud Cost Management Azure config",
	Long: `Delete Cloud Cost Management Azure config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#delete-cost-azure-uc-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err = api.DeleteCostAzureUCConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to delete-cost-azure-uc-config")

	},
}

func init() {

	Cmd.AddCommand(DeleteCostAzureUCConfigCmd)
}
