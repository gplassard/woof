package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:     "update-csm-threats-agent-policy [policy_id]",
	Aliases: []string{"update-agent-policy"},
	Short:   "Update a Workload Protection policy",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.UpdateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CloudWorkloadSecurityAgentPolicyUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-csm-threats-agent-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policy"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCSMThreatsAgentPolicyCmd)
}
