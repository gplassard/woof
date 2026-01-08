package csm_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use: "create-cloud-workload-security-agent-rule",

	Short: "Create a Workload Protection agent rule (US1-FED)",
	Long: `Create a Workload Protection agent rule (US1-FED)
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-threats/#create-cloud-workload-security-agent-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentRuleResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cloud-workload-security-agent-rule")

		fmt.Println(cmdutil.FormatJSON(res, "cloud_workload_security_agent_rule"))
	},
}

func init() {

	CreateCloudWorkloadSecurityAgentRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCloudWorkloadSecurityAgentRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCloudWorkloadSecurityAgentRuleCmd)
}
