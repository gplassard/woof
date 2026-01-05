package on_call

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:     "create-on-call-escalation-policy",
	Aliases: []string{"create-escalation-policy"},
	Short:   "Create On-Call escalation policy",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), datadogV2.EscalationPolicyCreateRequest{})
		cmdutil.HandleError(err, "failed to create-on-call-escalation-policy")

		cmd.Println(cmdutil.FormatJSON(res, "policies"))
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallEscalationPolicyCmd)
}
