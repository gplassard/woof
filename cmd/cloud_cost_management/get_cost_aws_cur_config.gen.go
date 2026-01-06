package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var GetCostAWSCURConfigCmd = &cobra.Command{
	Use: "get-cost-aws-cur-config [cloud_account_id]",

	Short: "Get cost AWS CUR config",
	Long: `Get cost AWS CUR config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#get-cost-aws-cur-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsCurConfigResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCostAWSCURConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-cost-aws-cur-config")

		cmd.Println(cmdutil.FormatJSON(res, "aws_cur_config"))
	},
}

func init() {

	Cmd.AddCommand(GetCostAWSCURConfigCmd)
}
