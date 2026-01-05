package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use: "create-cloud-workload-security-agent-rule [payload]",

	Short: "Create a Workload Protection agent rule (US1-FED)",
	Long: `Create a Workload Protection agent rule (US1-FED)
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#create-cloud-workload-security-agent-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRuleResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err = api.CreateCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cloud-workload-security-agent-rule")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateCloudWorkloadSecurityAgentRuleCmd)
}
