package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var DeleteCostAWSCURConfigCmd = &cobra.Command{
	Use: "delete-cost-aws-cur-config [cloud_account_id]",

	Short: "Delete Cloud Cost Management AWS CUR config",
	Long: `Delete Cloud Cost Management AWS CUR config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#delete-cost-aws-cur-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err = api.DeleteCostAWSCURConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to delete-cost-aws-cur-config")

	},
}

func init() {
	Cmd.AddCommand(DeleteCostAWSCURConfigCmd)
}
