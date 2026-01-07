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
	Long: `Get all CSM Serverless Agents
Documentation: https://docs.datadoghq.com/api/latest/csm-agents/#list-all-csm-serverless-agents`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CsmAgentsResponse
		var err error

		optionalParams := datadogV2.NewListAllCSMServerlessAgentsOptionalParameters()

		if cmd.Flags().Changed("page") {
			val, _ := cmd.Flags().GetInt64("page")
			optionalParams.WithPage(val)
		}

		if cmd.Flags().Changed("size") {
			val, _ := cmd.Flags().GetInt64("size")
			optionalParams.WithSize(val)
		}

		if cmd.Flags().Changed("query") {
			val, _ := cmd.Flags().GetString("query")
			optionalParams.WithQuery(val)
		}

		if cmd.Flags().Changed("order-direction") {
			val, _ := cmd.Flags().GetString("order-direction")
			optionalParams.WithOrderDirection(val)
		}

		api := datadogV2.NewCSMAgentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAllCSMServerlessAgents(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-all-csm-serverless-agents")

		cmd.Println(cmdutil.FormatJSON(res, "datadog_agent"))
	},
}

func init() {

	ListAllCSMServerlessAgentsCmd.Flags().Int64("page", 0, "The page index for pagination (zero-based).")

	ListAllCSMServerlessAgentsCmd.Flags().Int64("size", 0, "The number of items to include in a single page.")

	ListAllCSMServerlessAgentsCmd.Flags().String("query", "", "A search query string to filter results (for example, 'hostname:COMP-T2H4J27423').")

	ListAllCSMServerlessAgentsCmd.Flags().String("order-direction", "", "The sort direction for results. Use 'asc' for ascending or 'desc' for descending.")

	Cmd.AddCommand(ListAllCSMServerlessAgentsCmd)
}
