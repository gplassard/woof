package csm_threats

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use: "update-cloud-workload-security-agent-rule [agent_rule_id]",

	Short: "Update a Workload Protection agent rule (US1-FED)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.UpdateCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CloudWorkloadSecurityAgentRuleUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-cloud-workload-security-agent-rule")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCloudWorkloadSecurityAgentRuleCmd)
}
