package csm_threats

import (
	"fmt"
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
	Long: `Update a Workload Protection policy
Documentation: https://docs.datadoghq.com/api/latest/csm-threats/#update-csm-threats-agent-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudWorkloadSecurityAgentPolicyResponse
		var err error

		var body datadogV2.CloudWorkloadSecurityAgentPolicyUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCSMThreatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCSMThreatsAgentPolicy(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-csm-threats-agent-policy")

		fmt.Println(cmdutil.FormatJSON(res, "policy"))
	},
}

func init() {

	UpdateCSMThreatsAgentPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCSMThreatsAgentPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCSMThreatsAgentPolicyCmd)
}
