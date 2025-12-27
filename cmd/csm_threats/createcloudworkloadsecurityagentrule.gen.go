package csm_threats

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use:   "createcloudworkloadsecurityagentrule",
	Short: "Create a Workload Protection agent rule (US1-FED)",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCloudWorkloadSecurityAgentRule(client.NewContext(apiKey, appKey, site), datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createcloudworkloadsecurityagentrule: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_threats")
	},
}

func init() {
	Cmd.AddCommand(CreateCloudWorkloadSecurityAgentRuleCmd)
}
