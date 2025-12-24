package csm_threats

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "deletecsmthreatsagentpolicy [policy_id]",
	Short: "Delete a Workload Protection policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		_, err := api.DeleteCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletecsmthreatsagentpolicy: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCSMThreatsAgentPolicyCmd)
}
