package csm_threats

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCSMThreatsAgentRulesCmd = &cobra.Command{
	Use:   "list-csm-threats-agent-rules",
	Aliases: []string{ "list-agent-rules", },
	Short: "Get all Workload Protection agent rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.ListCSMThreatsAgentRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-csm-threats-agent-rules")

		cmdutil.PrintJSON(res, "agent_rule")
	},
}

func init() {
	Cmd.AddCommand(ListCSMThreatsAgentRulesCmd)
}
