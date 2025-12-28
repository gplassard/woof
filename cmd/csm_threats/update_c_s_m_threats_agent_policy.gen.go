package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "update_c_s_m_threats_agent_policy [policy_id]",
	Short: "Update a Workload Protection policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.UpdateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CloudWorkloadSecurityAgentPolicyUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_c_s_m_threats_agent_policy: %v", err)
		}

		cmdutil.PrintJSON(res, "policy")
	},
}

func init() {
	Cmd.AddCommand(UpdateCSMThreatsAgentPolicyCmd)
}
