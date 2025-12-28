package fleet_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListFleetAgentVersionsCmd = &cobra.Command{
	Use:   "list-fleet-agent-versions",
	
	Short: "List all available Agent versions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.ListFleetAgentVersions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-fleet-agent-versions: %v", err)
		}

		cmdutil.PrintJSON(res, "agent_version")
	},
}

func init() {
	Cmd.AddCommand(ListFleetAgentVersionsCmd)
}
