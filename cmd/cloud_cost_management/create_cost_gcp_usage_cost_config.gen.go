package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use: "create-cost-gcp-usage-cost-config [payload]",

	Short: "Create Google Cloud Usage Cost config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPUsageCostConfigResponse
		var err error

		var body datadogV2.GCPUsageCostConfigPostRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.CreateCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cost-gcp-usage-cost-config")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_uc_config"))
	},
}

func init() {
	Cmd.AddCommand(CreateCostGCPUsageCostConfigCmd)
}
