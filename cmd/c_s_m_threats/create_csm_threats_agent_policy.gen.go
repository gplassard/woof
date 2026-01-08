package c_s_m_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use: "create-csm-threats-agent-policy",

	Short: "Create a Workload Protection policy",
	Long: `Create a Workload Protection policy
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-threats/#create-csm-threats-agent-policy`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentPolicyResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentPolicyCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-csm-threats-agent-policy")

		fmt.Println(cmdutil.FormatJSON(res, "c_s_m_threats_agent_policy"))
	},
}

func init() {

	CreateCSMThreatsAgentPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCSMThreatsAgentPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCSMThreatsAgentPolicyCmd)
}
