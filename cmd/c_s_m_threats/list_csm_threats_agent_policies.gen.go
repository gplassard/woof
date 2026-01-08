package c_s_m_threats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCSMThreatsAgentPoliciesCmd = &cobra.Command{
	Use: "list-csm-threats-agent-policies",

	Short: "Get all Workload Protection policies",
	Long: `Get all Workload Protection policies
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-threats/#list-csm-threats-agent-policies`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentPoliciesListResponse
		var err error

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCSMThreatsAgentPolicies(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-csm-threats-agent-policies")

		fmt.Println(cmdutil.FormatJSON(res, "c_s_m_threats_agent_policie"))
	},
}

func init() {

	Cmd.AddCommand(ListCSMThreatsAgentPoliciesCmd)
}
