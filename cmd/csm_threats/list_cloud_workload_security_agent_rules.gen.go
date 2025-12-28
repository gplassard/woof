package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCloudWorkloadSecurityAgentRulesCmd = &cobra.Command{
	Use:   "list-cloud-workload-security-agent-rules",
	
	Short: "Get all Workload Protection agent rules (US1-FED)",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.ListCloudWorkloadSecurityAgentRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-cloud-workload-security-agent-rules: %v", err)
		}

		cmdutil.PrintJSON(res, "agent_rule")
	},
}

func init() {
	Cmd.AddCommand(ListCloudWorkloadSecurityAgentRulesCmd)
}
