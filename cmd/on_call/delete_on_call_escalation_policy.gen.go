package on_call

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "delete-on-call-escalation-policy [policy_id]",
	Short: "Delete On-Call escalation policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		_, err := api.DeleteOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-on-call-escalation-policy: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteOnCallEscalationPolicyCmd)
}
