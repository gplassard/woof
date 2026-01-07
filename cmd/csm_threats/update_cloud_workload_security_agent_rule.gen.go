package csm_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use: "update-cloud-workload-security-agent-rule [agent_rule_id]",

	Short: "Update a Workload Protection agent rule (US1-FED)",
	Long: `Update a Workload Protection agent rule (US1-FED)
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#update-cloud-workload-security-agent-rule`,
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
		res, _, err = api.UpdateCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-cloud-workload-security-agent-rule")

		fmt.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {

	UpdateCloudWorkloadSecurityAgentRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCloudWorkloadSecurityAgentRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCloudWorkloadSecurityAgentRuleCmd)
}
