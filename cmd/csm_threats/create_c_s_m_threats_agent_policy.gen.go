package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "create_c_s_m_threats_agent_policy",
	Short: "Create a Workload Protection policy",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), datadogV2.CloudWorkloadSecurityAgentPolicyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_c_s_m_threats_agent_policy: %v", err)
		}

		cmdutil.PrintJSON(res, "policy")
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentPolicyCmd)
}
