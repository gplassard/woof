package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCostAzureUCConfigsCmd = &cobra.Command{
	Use: "create-cost-azure-uc-configs [payload]",

	Short: "Create Cloud Cost Management Azure configs",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureUCConfigPairsResponse
		var err error

		var body datadogV2.AzureUCConfigPostRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.CreateCostAzureUCConfigs(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cost-azure-uc-configs")

		cmd.Println(cmdutil.FormatJSON(res, "azure_uc_configs"))
	},
}

func init() {
	Cmd.AddCommand(CreateCostAzureUCConfigsCmd)
}
