package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "get-c-s-m-threats-agent-policy [policy_id]",
	Short: "Get a Workload Protection policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.GetCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-c-s-m-threats-agent-policy: %v", err)
		}

		cmdutil.PrintJSON(res, "policy")
	},
}

func init() {
	Cmd.AddCommand(GetCSMThreatsAgentPolicyCmd)
}
