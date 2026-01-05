package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:     "get-csm-threats-agent-policy [policy_id]",
	Aliases: []string{"get-agent-policy"},
	Short:   "Get a Workload Protection policy",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentPolicyResponse
		var err error

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err = api.GetCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-csm-threats-agent-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policy"))
	},
}

func init() {
	Cmd.AddCommand(GetCSMThreatsAgentPolicyCmd)
}
