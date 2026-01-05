package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCostAWSCURConfigCmd = &cobra.Command{
	Use: "create-cost-aws-cur-config [payload]",

	Short: "Create Cloud Cost Management AWS CUR config",
	Long: `Create Cloud Cost Management AWS CUR config
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#create-cost-aws-cur-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsCurConfigResponse
		var err error

		var body datadogV2.AwsCURConfigPostRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.CreateCostAWSCURConfig(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cost-aws-cur-config")

		cmd.Println(cmdutil.FormatJSON(res, "aws_cur_config"))
	},
}

func init() {
	Cmd.AddCommand(CreateCostAWSCURConfigCmd)
}
