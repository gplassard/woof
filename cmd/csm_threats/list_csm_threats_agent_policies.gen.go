package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCSMThreatsAgentPoliciesCmd = &cobra.Command{
	Use:   "list-csm-threats-agent-policies",
	Aliases: []string{ "list-agent-policies", },
	Short: "Get all Workload Protection policies",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.ListCSMThreatsAgentPolicies(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-csm-threats-agent-policies: %v", err)
		}

		cmdutil.PrintJSON(res, "policy")
	},
}

func init() {
	Cmd.AddCommand(ListCSMThreatsAgentPoliciesCmd)
}
