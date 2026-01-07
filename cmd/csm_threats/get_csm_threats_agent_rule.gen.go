package csm_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:     "get-csm-threats-agent-rule [agent_rule_id]",
	Aliases: []string{"get-agent-rule"},
	Short:   "Get a Workload Protection agent rule",
	Long: `Get a Workload Protection agent rule
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#get-csm-threats-agent-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRuleResponse
		var err error

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-csm-threats-agent-rule")

		fmt.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetCSMThreatsAgentRuleCmd)
}
