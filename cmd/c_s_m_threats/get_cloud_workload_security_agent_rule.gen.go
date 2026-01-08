package c_s_m_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use: "get-cloud-workload-security-agent-rule [agent_rule_id]",

	Short: "Get a Workload Protection agent rule (US1-FED)",
	Long: `Get a Workload Protection agent rule (US1-FED)
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-threats/#get-cloud-workload-security-agent-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRuleResponse
		var err error

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-cloud-workload-security-agent-rule")

		fmt.Println(cmdutil.FormatJSON(res, "cloud_workload_security_agent_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetCloudWorkloadSecurityAgentRuleCmd)
}
