package csm_threats

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCSMThreatsAgentRulesCmd = &cobra.Command{
	Use:   "listcsmthreatsagentrules",
	Short: "Get all Workload Protection agent rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.ListCSMThreatsAgentRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listcsmthreatsagentrules: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_threats")
	},
}

func init() {
	Cmd.AddCommand(ListCSMThreatsAgentRulesCmd)
}
