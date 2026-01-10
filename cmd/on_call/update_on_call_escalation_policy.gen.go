package on_call

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:     "update-on-call-escalation-policy [policy_id]",
	Aliases: []string{"update-escalation-policy"},
	Short:   "Update On-Call escalation policy",
	Long: `Update On-Call escalation policy
Documentation: https://docs.datadoghq.com/api/latest/on-call/#update-on-call-escalation-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EscalationPolicy
		var err error

		optionalParams := datadogV2.NewUpdateOnCallEscalationPolicyOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to update-on-call-escalation-policy")

		fmt.Println(cmdutil.FormatJSON(res, "policies"))
	},
}

func init() {

	UpdateOnCallEscalationPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateOnCallEscalationPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	UpdateOnCallEscalationPolicyCmd.Flags().String("include", "", "Comma-separated list of included relationships to be returned. Allowed values: 'teams', 'steps', 'steps.targets'.")

	Cmd.AddCommand(UpdateOnCallEscalationPolicyCmd)
}
