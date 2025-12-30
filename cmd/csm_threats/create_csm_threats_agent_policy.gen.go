package csm_threats

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:     "create-csm-threats-agent-policy",
	Aliases: []string{"create-agent-policy"},
	Short:   "Create a Workload Protection policy",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), datadogV2.CloudWorkloadSecurityAgentPolicyCreateRequest{})
		cmdutil.HandleError(err, "failed to create-csm-threats-agent-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policy"))
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentPolicyCmd)
}
