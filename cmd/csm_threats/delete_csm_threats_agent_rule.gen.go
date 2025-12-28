package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "delete-csm-threats-agent-rule [agent_rule_id]",
	Aliases: []string{ "delete-agent-rule", },
	Short: "Delete a Workload Protection agent rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		_, err := api.DeleteCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-csm-threats-agent-rule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCSMThreatsAgentRuleCmd)
}
