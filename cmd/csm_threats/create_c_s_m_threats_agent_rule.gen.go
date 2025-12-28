package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "create-c-s-m-threats-agent-rule",
	
	Short: "Create a Workload Protection agent rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-c-s-m-threats-agent-rule: %v", err)
		}

		cmdutil.PrintJSON(res, "agent_rule")
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentRuleCmd)
}
