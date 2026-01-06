package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFleetAgentsCmd = &cobra.Command{
	Use: "list-fleet-agents",

	Short: "List all Datadog Agents",
	Long: `List all Datadog Agents
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#list-fleet-agents`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetAgentsResponse
		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFleetAgents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fleet-agents")

		cmd.Println(cmdutil.FormatJSON(res, "fleet_automation"))
	},
}

func init() {

	Cmd.AddCommand(ListFleetAgentsCmd)
}
