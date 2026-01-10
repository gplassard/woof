package csm_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCSMThreatsAgentRulesCmd = &cobra.Command{
	Use:     "list-csm-threats-agent-rules",
	Aliases: []string{"list-agent-rules"},
	Short:   "Get all Workload Protection agent rules",
	Long: `Get all Workload Protection agent rules
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#list-csm-threats-agent-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRulesListResponse
		var err error

		optionalParams := datadogV2.NewListCSMThreatsAgentRulesOptionalParameters()

		if cmd.Flags().Changed("policy-id") {
			val, _ := cmd.Flags().GetString("policy-id")
			optionalParams.WithPolicyId(val)
		}

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCSMThreatsAgentRules(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-csm-threats-agent-rules")

		fmt.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {

	ListCSMThreatsAgentRulesCmd.Flags().String("policy-id", "", "The ID of the Agent policy")

	Cmd.AddCommand(ListCSMThreatsAgentRulesCmd)
}
