package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCostAzureUCConfigsCmd = &cobra.Command{
	Use: "create-cost-azure-uc-configs",

	Short: "Create Cloud Cost Management Azure configs",
	Long: `Create Cloud Cost Management Azure configs
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#create-cost-azure-uc-configs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureUCConfigPairsResponse
		var err error

		var body datadogV2.AzureUCConfigPostRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCostAzureUCConfigs(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cost-azure-uc-configs")

		fmt.Println(cmdutil.FormatJSON(res, "azure_uc_configs"))
	},
}

func init() {

	CreateCostAzureUCConfigsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCostAzureUCConfigsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCostAzureUCConfigsCmd)
}
