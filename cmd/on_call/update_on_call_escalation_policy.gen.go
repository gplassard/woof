package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:     "update-on-call-escalation-policy [policy_id] [payload]",
	Aliases: []string{"update-escalation-policy"},
	Short:   "Update On-Call escalation policy",
	Long: `Update On-Call escalation policy
Documentation: https://docs.datadoghq.com/api/latest/on-call/#update-on-call-escalation-policy`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EscalationPolicy
		var err error

		var body datadogV2.EscalationPolicyUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err = api.UpdateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-on-call-escalation-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policies"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallEscalationPolicyCmd)
}
