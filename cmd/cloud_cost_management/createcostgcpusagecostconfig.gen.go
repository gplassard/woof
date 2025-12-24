package cloud_cost_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "createcostgcpusagecostconfig",
	Short: "Create Google Cloud Usage Cost config",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCostGCPUsageCostConfig(client.NewContext(apiKey, appKey, site), datadogV2.GCPUsageCostConfigPostRequest{})
		if err != nil {
			log.Fatalf("failed to createcostgcpusagecostconfig: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(CreateCostGCPUsageCostConfigCmd)
}
