package cloud_cost_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	"strconv"
)

var DeleteCustomAllocationRuleCmd = &cobra.Command{
	Use:   "deletecustomallocationrule [rule_id]",
	Short: "Delete custom allocation rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.DeleteCustomAllocationRule(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		if err != nil {
			log.Fatalf("failed to deletecustomallocationrule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomAllocationRuleCmd)
}
