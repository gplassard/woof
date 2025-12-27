package csm_agents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAllCSMServerlessAgentsCmd = &cobra.Command{
	Use:   "listallcsmserverlessagents",
	Short: "Get all CSM Serverless Agents",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMAgentsApi(client.NewAPIClient())
		res, _, err := api.ListAllCSMServerlessAgents(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listallcsmserverlessagents: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_agents")
	},
}

func init() {
	Cmd.AddCommand(ListAllCSMServerlessAgentsCmd)
}
