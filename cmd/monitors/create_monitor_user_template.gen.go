package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateMonitorUserTemplateCmd = &cobra.Command{
	Use:   "create_monitor_user_template",
	Short: "Create a monitor user template",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.CreateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), datadogV2.MonitorUserTemplateCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_monitor_user_template: %v", err)
		}

		cmdutil.PrintJSON(res, "monitors")
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorUserTemplateCmd)
}
