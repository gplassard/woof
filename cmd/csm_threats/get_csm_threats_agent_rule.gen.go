package csm_threats

import (
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

		optionalParams := datadogV2.NewGetCSMThreatsAgentRuleOptionalParameters()

		if cmd.Flags().Changed("policy-id") {
			val, _ := cmd.Flags().GetString("policy-id")
			optionalParams.WithPolicyId(val)
		}

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-csm-threats-agent-rule")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {

	GetCSMThreatsAgentRuleCmd.Flags().String("policy-id", "", "The ID of the Agent policy")

	Cmd.AddCommand(GetCSMThreatsAgentRuleCmd)
}
