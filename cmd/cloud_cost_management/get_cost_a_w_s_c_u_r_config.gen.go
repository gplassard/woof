package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var GetCostAWSCURConfigCmd = &cobra.Command{
	Use:   "get_cost_a_w_s_c_u_r_config [cloud_account_id]",
	Short: "Get cost AWS CUR config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.GetCostAWSCURConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		if err != nil {
			log.Fatalf("failed to get_cost_a_w_s_c_u_r_config: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_cur_config")
	},
}

func init() {
	Cmd.AddCommand(GetCostAWSCURConfigCmd)
}
