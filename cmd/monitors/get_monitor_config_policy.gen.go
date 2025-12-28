package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "get-monitor-config-policy [policy_id]",
	Short: "Get a monitor configuration policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.GetMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-monitor-config-policy: %v", err)
		}

		cmdutil.PrintJSON(res, "monitor-config-policy")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorConfigPolicyCmd)
}
