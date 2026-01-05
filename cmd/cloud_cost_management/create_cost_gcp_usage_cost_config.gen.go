package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use: "create-cost-gcp-usage-cost-config",

	Short: "Create Google Cloud Usage Cost config",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), datadogV2.GCPUsageCostConfigPostRequest{})
		cmdutil.HandleError(err, "failed to create-cost-gcp-usage-cost-config")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_uc_config"))
	},
}

func init() {
	Cmd.AddCommand(CreateCostGCPUsageCostConfigCmd)
}
