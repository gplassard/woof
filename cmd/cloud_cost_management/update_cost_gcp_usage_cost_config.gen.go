package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var UpdateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use: "update-cost-gcp-usage-cost-config [cloud_account_id]",

	Short: "Update Google Cloud Usage Cost config",
	Long: `Update Google Cloud Usage Cost config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#update-cost-gcp-usage-cost-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPUsageCostConfigResponse
		var err error

		var body datadogV2.GCPUsageCostConfigPatchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to update-cost-gcp-usage-cost-config")

		fmt.Println(cmdutil.FormatJSON(res, "cost_g_c_p_usage_cost_config"))
	},
}

func init() {

	UpdateCostGCPUsageCostConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCostGCPUsageCostConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCostGCPUsageCostConfigCmd)
}
