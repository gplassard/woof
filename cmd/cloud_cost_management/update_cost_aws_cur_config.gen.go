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

var UpdateCostAWSCURConfigCmd = &cobra.Command{
	Use: "update-cost-aws-cur-config [cloud_account_id]",

	Short: "Update Cloud Cost Management AWS CUR config",
	Long: `Update Cloud Cost Management AWS CUR config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#update-cost-aws-cur-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsCURConfigsResponse
		var err error

		var body datadogV2.AwsCURConfigPatchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCostAWSCURConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to update-cost-aws-cur-config")

		fmt.Println(cmdutil.FormatJSON(res, "aws_cur_config"))
	},
}

func init() {

	UpdateCostAWSCURConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCostAWSCURConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCostAWSCURConfigCmd)
}
