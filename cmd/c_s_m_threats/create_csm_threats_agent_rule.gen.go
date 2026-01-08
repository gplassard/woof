package c_s_m_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCSMThreatsAgentRuleCmd = &cobra.Command{
	Use: "create-csm-threats-agent-rule",

	Short: "Create a Workload Protection agent rule",
	Long: `Create a Workload Protection agent rule
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-threats/#create-csm-threats-agent-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRuleResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-csm-threats-agent-rule")

		fmt.Println(cmdutil.FormatJSON(res, "c_s_m_threats_agent_rule"))
	},
}

func init() {

	CreateCSMThreatsAgentRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCSMThreatsAgentRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCSMThreatsAgentRuleCmd)
}
