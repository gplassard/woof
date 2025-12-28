package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "get-csm-threats-agent-rule [agent_rule_id]",
	Aliases: []string{ "get-agent-rule", },
	Short: "Get a Workload Protection agent rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.GetCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-csm-threats-agent-rule: %v", err)
		}

		cmdutil.PrintJSON(res, "agent_rule")
	},
}

func init() {
	Cmd.AddCommand(GetCSMThreatsAgentRuleCmd)
}
