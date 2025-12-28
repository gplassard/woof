package on_call

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "get-on-call-escalation-policy [policy_id]",
	Aliases: []string{ "get-escalation-policy", },
	Short: "Get On-Call escalation policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.GetOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-on-call-escalation-policy")

		cmdutil.PrintJSON(res, "policies")
	},
}

func init() {
	Cmd.AddCommand(GetOnCallEscalationPolicyCmd)
}
