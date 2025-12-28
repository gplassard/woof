package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListMonitorUserTemplatesCmd = &cobra.Command{
	Use:   "list_monitor_user_templates",
	Short: "Get all monitor user templates",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.ListMonitorUserTemplates(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_monitor_user_templates: %v", err)
		}

		cmdutil.PrintJSON(res, "monitor-user-template")
	},
}

func init() {
	Cmd.AddCommand(ListMonitorUserTemplatesCmd)
}
