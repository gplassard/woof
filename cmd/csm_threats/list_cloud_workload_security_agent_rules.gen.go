package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCloudWorkloadSecurityAgentRulesCmd = &cobra.Command{
	Use: "list-cloud-workload-security-agent-rules",

	Short: "Get all Workload Protection agent rules (US1-FED)",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.ListCloudWorkloadSecurityAgentRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cloud-workload-security-agent-rules")

		cmd.Println(cmdutil.FormatJSON(res, "agent_rule"))
	},
}

func init() {
	Cmd.AddCommand(ListCloudWorkloadSecurityAgentRulesCmd)
}
