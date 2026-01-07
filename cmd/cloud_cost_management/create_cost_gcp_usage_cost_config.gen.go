package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use: "create-cost-gcp-usage-cost-config",

	Short: "Create Google Cloud Usage Cost config",
	Long: `Create Google Cloud Usage Cost config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#create-cost-gcp-usage-cost-config`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPUsageCostConfigResponse
		var err error

		var body datadogV2.GCPUsageCostConfigPostRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cost-gcp-usage-cost-config")

		fmt.Println(cmdutil.FormatJSON(res, "gcp_uc_config"))
	},
}

func init() {

	CreateCostGCPUsageCostConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCostGCPUsageCostConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCostGCPUsageCostConfigCmd)
}
