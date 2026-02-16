package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:     "delete-csm-threats-agent-rule [agent_rule_id]",
	Aliases: []string{"delete-agent-rule"},
	Short:   "Delete a Workload Protection agent rule",
	Long: `Delete a Workload Protection agent rule
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#delete-csm-threats-agent-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		optionalParams := datadogV2.NewDeleteCSMThreatsAgentRuleOptionalParameters()

		if cmd.Flags().Changed("policy-id") {
			val, _ := cmd.Flags().GetString("policy-id")
			optionalParams.WithPolicyId(val)
		}

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to delete-csm-threats-agent-rule")

	},
}

func init() {

	DeleteCSMThreatsAgentRuleCmd.Flags().String("policy-id", "", "The ID of the Agent policy")

	Cmd.AddCommand(DeleteCSMThreatsAgentRuleCmd)
}
