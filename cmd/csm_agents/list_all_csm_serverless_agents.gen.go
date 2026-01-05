package csm_agents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAllCSMServerlessAgentsCmd = &cobra.Command{
	Use: "list-all-csm-serverless-agents",

	Short: "Get all CSM Serverless Agents",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMAgentsApi(client.NewAPIClient())
		res, _, err := api.ListAllCSMServerlessAgents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-all-csm-serverless-agents")

		cmd.Println(cmdutil.FormatJSON(res, "datadog_agent"))
	},
}

func init() {
	Cmd.AddCommand(ListAllCSMServerlessAgentsCmd)
}
