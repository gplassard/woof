package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetOnCallEscalationPolicyCmd = &cobra.Command{
	Use:     "get-on-call-escalation-policy [policy_id]",
	Aliases: []string{"get-escalation-policy"},
	Short:   "Get On-Call escalation policy",
	Long: `Get On-Call escalation policy
Documentation: https://docs.datadoghq.com/api/latest/on-call/#get-on-call-escalation-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EscalationPolicy
		var err error

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.GetOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-on-call-escalation-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policies"))
	},
}

func init() {

	Cmd.AddCommand(GetOnCallEscalationPolicyCmd)
}
