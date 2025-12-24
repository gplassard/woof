package csm_threats

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "createcsmthreatsagentpolicy",
	Short: "Create a Workload Protection policy",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), datadogV2.CloudWorkloadSecurityAgentPolicyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createcsmthreatsagentpolicy: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_threats")
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentPolicyCmd)
}
