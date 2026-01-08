package on_call

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:     "create-on-call-escalation-policy",
	Aliases: []string{"create-escalation-policy"},
	Short:   "Create On-Call escalation policy",
	Long: `Create On-Call escalation policy
Documentation: https://docs.datadoghq.com/api/latest/on-call/#create-on-call-escalation-policy`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EscalationPolicy
		var err error

		var body datadogV2.EscalationPolicyCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-on-call-escalation-policy")

		fmt.Println(cmdutil.FormatJSON(res, "on_call_escalation_policy"))
	},
}

func init() {

	CreateOnCallEscalationPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateOnCallEscalationPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateOnCallEscalationPolicyCmd)
}
