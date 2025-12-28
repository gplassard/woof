package csm_agents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAllCSMAgentsCmd = &cobra.Command{
	Use:   "list-all-c-s-m-agents",
	Short: "Get all CSM Agents",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMAgentsApi(client.NewAPIClient())
		res, _, err := api.ListAllCSMAgents(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-all-c-s-m-agents: %v", err)
		}

		cmdutil.PrintJSON(res, "datadog_agent")
	},
}

func init() {
	Cmd.AddCommand(ListAllCSMAgentsCmd)
}
