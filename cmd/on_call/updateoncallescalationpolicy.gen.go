package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "updateoncallescalationpolicy [policy_id]",
	Short: "Update On-Call escalation policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.UpdateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0], datadogV2.EscalationPolicyUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateoncallescalationpolicy: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call")
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallEscalationPolicyCmd)
}
