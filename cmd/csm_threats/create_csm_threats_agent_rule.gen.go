package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:     "create-csm-threats-agent-rule",
	Aliases: []string{"create-agent-rule"},
	Short:   "Create a Workload Protection agent rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-csm-threats-agent-rule")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentRuleCmd)
}
