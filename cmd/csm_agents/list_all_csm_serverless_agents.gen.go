package csm_agents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAllCSMServerlessAgentsCmd = &cobra.Command{
	Use: "list-all-csm-serverless-agents",

	Short: "Get all CSM Serverless Agents",
	Long: `Get all CSM Serverless Agents
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-agents/#list-all-csm-serverless-agents`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CsmAgentsResponse
		var err error

		api := datadogV2.NewCSMAgentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAllCSMServerlessAgents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-all-csm-serverless-agents")

		fmt.Println(cmdutil.FormatJSON(res, "all_c_s_m_serverless_agent"))
	},
}

func init() {

	Cmd.AddCommand(ListAllCSMServerlessAgentsCmd)
}
