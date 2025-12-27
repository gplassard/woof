package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "create_on_call_escalation_policy",
	Short: "Create On-Call escalation policy",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), datadogV2.EscalationPolicyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_on_call_escalation_policy: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallEscalationPolicyCmd)
}
