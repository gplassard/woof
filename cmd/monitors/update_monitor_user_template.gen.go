package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateMonitorUserTemplateCmd = &cobra.Command{
	Use:   "update-monitor-user-template [template_id]",
	Short: "Update a monitor user template to a new version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.UpdateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MonitorUserTemplateUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-monitor-user-template: %v", err)
		}

		cmdutil.PrintJSON(res, "monitor-user-template")
	},
}

func init() {
	Cmd.AddCommand(UpdateMonitorUserTemplateCmd)
}
