package csm_threats

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "update-csm-threats-agent-rule [agent_rule_id]",
	Aliases: []string{ "update-agent-rule", },
	Short: "Update a Workload Protection agent rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.UpdateCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CloudWorkloadSecurityAgentRuleUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-csm-threats-agent-rule")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCSMThreatsAgentRuleCmd)
}
