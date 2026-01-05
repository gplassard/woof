package on_call

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteOnCallEscalationPolicyCmd = &cobra.Command{
	Use:     "delete-on-call-escalation-policy [policy_id]",
	Aliases: []string{"delete-escalation-policy"},
	Short:   "Delete On-Call escalation policy",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		_, err := api.DeleteOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-on-call-escalation-policy")

	},
}

func init() {
	Cmd.AddCommand(DeleteOnCallEscalationPolicyCmd)
}
