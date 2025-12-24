package csm_agents

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAllCSMAgentsCmd = &cobra.Command{
	Use:   "listallcsmagents",
	Short: "Get all CSM Agents",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCSMAgentsApi(client.NewAPIClient())
		res, _, err := api.ListAllCSMAgents(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listallcsmagents: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_agents")
	},
}

func init() {
	Cmd.AddCommand(ListAllCSMAgentsCmd)
}
