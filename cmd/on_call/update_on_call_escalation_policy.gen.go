package on_call

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "update-on-call-escalation-policy [policy_id]",
	Aliases: []string{ "update-escalation-policy", },
	Short: "Update On-Call escalation policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.UpdateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0], datadogV2.EscalationPolicyUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-on-call-escalation-policy")

		cmdutil.PrintJSON(res, "policies")
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallEscalationPolicyCmd)
}
