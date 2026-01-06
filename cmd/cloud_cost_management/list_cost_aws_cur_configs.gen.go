package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCostAWSCURConfigsCmd = &cobra.Command{
	Use: "list-cost-aws-cur-configs",

	Short: "List Cloud Cost Management AWS CUR configs",
	Long: `List Cloud Cost Management AWS CUR configs
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#list-cost-aws-cur-configs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsCURConfigsResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCostAWSCURConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cost-aws-cur-configs")

		cmd.Println(cmdutil.FormatJSON(res, "aws_cur_config"))
	},
}

func init() {

	Cmd.AddCommand(ListCostAWSCURConfigsCmd)
}
