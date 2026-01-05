package csm_threats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:     "create-csm-threats-agent-policy [payload]",
	Aliases: []string{"create-agent-policy"},
	Short:   "Create a Workload Protection policy",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.CloudWorkloadSecurityAgentPolicyCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		res, _, err := api.CreateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-csm-threats-agent-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policy"))
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentPolicyCmd)
}
