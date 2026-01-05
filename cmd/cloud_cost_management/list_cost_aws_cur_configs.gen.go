package cloud_cost_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCostAWSCURConfigsCmd = &cobra.Command{
	Use: "list-cost-aws-cur-configs",

	Short: "List Cloud Cost Management AWS CUR configs",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCostAWSCURConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cost-aws-cur-configs")

		cmd.Println(cmdutil.FormatJSON(res, "aws_cur_config"))
	},
}

func init() {
	Cmd.AddCommand(ListCostAWSCURConfigsCmd)
}
