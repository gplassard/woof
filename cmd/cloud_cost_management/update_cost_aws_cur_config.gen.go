package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var UpdateCostAWSCURConfigCmd = &cobra.Command{
	Use: "update-cost-aws-cur-config [cloud_account_id]",

	Short: "Update Cloud Cost Management AWS CUR config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCostAWSCURConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.AwsCURConfigPatchRequest{})
		cmdutil.HandleError(err, "failed to update-cost-aws-cur-config")

		cmd.Println(cmdutil.FormatJSON(res, "aws_cur_config"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCostAWSCURConfigCmd)
}
