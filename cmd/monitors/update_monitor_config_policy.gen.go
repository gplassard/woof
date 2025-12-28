package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "update-monitor-config-policy [policy_id]",
	Short: "Edit a monitor configuration policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.UpdateMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MonitorConfigPolicyEditRequest{})
		if err != nil {
			log.Fatalf("failed to update-monitor-config-policy: %v", err)
		}

		cmdutil.PrintJSON(res, "monitor-config-policy")
	},
}

func init() {
	Cmd.AddCommand(UpdateMonitorConfigPolicyCmd)
}
