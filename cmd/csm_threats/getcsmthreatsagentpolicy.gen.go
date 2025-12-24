package csm_threats

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "getcsmthreatsagentpolicy [policy_id]",
	Short: "Get a Workload Protection policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.GetCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getcsmthreatsagentpolicy: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_threats")
	},
}

func init() {
	Cmd.AddCommand(GetCSMThreatsAgentPolicyCmd)
}
