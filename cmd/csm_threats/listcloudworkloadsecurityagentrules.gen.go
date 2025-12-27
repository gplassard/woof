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
	Use:   "listcloudworkloadsecurityagentrules",
	Short: "Get all Workload Protection agent rules (US1-FED)",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.ListCloudWorkloadSecurityAgentRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listcloudworkloadsecurityagentrules: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_threats")
	},
}

func init() {
	Cmd.AddCommand(ListCloudWorkloadSecurityAgentRulesCmd)
}
