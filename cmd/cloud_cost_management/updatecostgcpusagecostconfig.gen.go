package cloud_cost_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var UpdateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "updatecostgcpusagecostconfig [cloud_account_id]",
	Short: "Update Google Cloud Usage Cost config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.GCPUsageCostConfigPatchRequest{})
		if err != nil {
			log.Fatalf("failed to updatecostgcpusagecostconfig: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(UpdateCostGCPUsageCostConfigCmd)
}
