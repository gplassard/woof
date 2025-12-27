package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "delete_monitor_config_policy [policy_id]",
	Short: "Delete a monitor configuration policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.DeleteMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_monitor_config_policy: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteMonitorConfigPolicyCmd)
}
