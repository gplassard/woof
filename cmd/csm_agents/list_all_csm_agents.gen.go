package csm_agents

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAllCSMAgentsCmd = &cobra.Command{
	Use:     "list-all-csm-agents",
	Aliases: []string{"list-all"},
	Short:   "Get all CSM Agents",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMAgentsApi(client.NewAPIClient())
		res, _, err := api.ListAllCSMAgents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-all-csm-agents")

		cmd.Println(cmdutil.FormatJSON(res, "datadog_agent"))
	},
}

func init() {
	Cmd.AddCommand(ListAllCSMAgentsCmd)
}
