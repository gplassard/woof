package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFleetAgentVersionsCmd = &cobra.Command{
	Use: "list-fleet-agent-versions",

	Short: "List all available Agent versions",
	Long: `List all available Agent versions
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#list-fleet-agent-versions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetAgentVersionsResponse
		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFleetAgentVersions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fleet-agent-versions")

		cmd.Println(cmdutil.FormatJSON(res, "agent_version"))
	},
}

func init() {

	Cmd.AddCommand(ListFleetAgentVersionsCmd)
}
