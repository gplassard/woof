package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCostAWSCURConfigCmd = &cobra.Command{
	Use: "create-cost-aws-cur-config",

	Short: "Create Cloud Cost Management AWS CUR config",
	Long: `Create Cloud Cost Management AWS CUR config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#create-cost-aws-cur-config`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsCurConfigResponse
		var err error

		var body datadogV2.AwsCURConfigPostRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.CreateCostAWSCURConfig(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cost-aws-cur-config")

		cmd.Println(cmdutil.FormatJSON(res, "aws_cur_config"))
	},
}

func init() {

	CreateCostAWSCURConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCostAWSCURConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCostAWSCURConfigCmd)
}
