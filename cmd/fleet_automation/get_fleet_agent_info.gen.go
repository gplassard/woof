package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFleetAgentInfoCmd = &cobra.Command{
	Use: "get-fleet-agent-info [agent_key]",

	Short: "Get detailed information about an agent",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.GetFleetAgentInfo(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-fleet-agent-info")

		cmd.Println(cmdutil.FormatJSON(res, "datadog_agent_key"))
	},
}

func init() {
	Cmd.AddCommand(GetFleetAgentInfoCmd)
}
