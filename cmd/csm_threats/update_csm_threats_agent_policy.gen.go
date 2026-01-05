package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:     "update-csm-threats-agent-policy [policy_id] [payload]",
	Aliases: []string{"update-agent-policy"},
	Short:   "Update a Workload Protection policy",
	Long: `Update a Workload Protection policy
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#update-csm-threats-agent-policy`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentPolicyResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentPolicyUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err = api.UpdateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-csm-threats-agent-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policy"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCSMThreatsAgentPolicyCmd)
}
