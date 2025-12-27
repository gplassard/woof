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
	Use:   "createoncallescalationpolicy",
	Short: "Create On-Call escalation policy",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallEscalationPolicy(client.NewContext(apiKey, appKey, site), datadogV2.EscalationPolicyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createoncallescalationpolicy: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallEscalationPolicyCmd)
}
