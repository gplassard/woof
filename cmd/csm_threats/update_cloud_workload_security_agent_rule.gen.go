package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use: "update-cloud-workload-security-agent-rule [agent_rule_id] [payload]",

	Short: "Update a Workload Protection agent rule (US1-FED)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRuleResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentRuleUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err = api.UpdateCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-cloud-workload-security-agent-rule")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCloudWorkloadSecurityAgentRuleCmd)
}
