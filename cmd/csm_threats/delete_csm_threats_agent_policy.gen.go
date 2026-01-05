package csm_threats

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:     "delete-csm-threats-agent-policy [policy_id]",
	Aliases: []string{"delete-agent-policy"},
	Short:   "Delete a Workload Protection policy",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		_, err := api.DeleteCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-csm-threats-agent-policy")

	},
}

func init() {
	Cmd.AddCommand(DeleteCSMThreatsAgentPolicyCmd)
}
