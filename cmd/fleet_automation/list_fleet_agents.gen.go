package fleet_automation

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFleetAgentsCmd = &cobra.Command{
	Use: "list-fleet-agents",

	Short: "List all Datadog Agents",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.ListFleetAgents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fleet-agents")

		cmd.Println(cmdutil.FormatJSON(res, "fleet_automation"))
	},
}

func init() {
	Cmd.AddCommand(ListFleetAgentsCmd)
}
