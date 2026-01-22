package csm_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:     "update-csm-threats-agent-rule [agent_rule_id]",
	Aliases: []string{"update-agent-rule"},
	Short:   "Update a Workload Protection agent rule",
	Long: `Update a Workload Protection agent rule
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#update-csm-threats-agent-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRuleResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentRuleUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCSMThreatsAgentRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-csm-threats-agent-rule")

		fmt.Println(cmdutil.FormatJSON(res, "c_s_m_threats_agent_rule"))
	},
}

func init() {

	UpdateCSMThreatsAgentRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCSMThreatsAgentRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCSMThreatsAgentRuleCmd)
}
