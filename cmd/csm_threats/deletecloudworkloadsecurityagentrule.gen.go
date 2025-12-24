package csm_threats

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use:   "deletecloudworkloadsecurityagentrule [agent_rule_id]",
	Short: "Delete a Workload Protection agent rule (US1-FED)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		_, err := api.DeleteCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletecloudworkloadsecurityagentrule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCloudWorkloadSecurityAgentRuleCmd)
}
