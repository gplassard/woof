package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCostAWSCURConfigCmd = &cobra.Command{
	Use:   "create_cost_a_w_s_c_u_r_config",
	Short: "Create Cloud Cost Management AWS CUR config",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCostAWSCURConfig(client.NewContext(apiKey, appKey, site), datadogV2.AwsCURConfigPostRequest{})
		if err != nil {
			log.Fatalf("failed to create_cost_a_w_s_c_u_r_config: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_cur_config")
	},
}

func init() {
	Cmd.AddCommand(CreateCostAWSCURConfigCmd)
}
