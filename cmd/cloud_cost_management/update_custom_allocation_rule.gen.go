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

var UpdateCustomAllocationRuleCmd = &cobra.Command{
	Use:   "update-custom-allocation-rule [rule_id]",
	Short: "Update custom allocation rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCustomAllocationRule(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.ArbitraryCostUpsertRequest{})
		if err != nil {
			log.Fatalf("failed to update-custom-allocation-rule: %v", err)
		}

		cmdutil.PrintJSON(res, "arbitrary_rule")
	},
}

func init() {
	Cmd.AddCommand(UpdateCustomAllocationRuleCmd)
}
