package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use: "create-cloud-workload-security-agent-rule",

	Short: "Create a Workload Protection agent rule (US1-FED)",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-cloud-workload-security-agent-rule")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateCloudWorkloadSecurityAgentRuleCmd)
}
