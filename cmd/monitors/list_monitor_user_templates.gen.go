package monitors

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListMonitorUserTemplatesCmd = &cobra.Command{
	Use:   "list-monitor-user-templates",
	Aliases: []string{ "list-user-templates", },
	Short: "Get all monitor user templates",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.ListMonitorUserTemplates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-monitor-user-templates")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-user-template"))
	},
}

func init() {
	Cmd.AddCommand(ListMonitorUserTemplatesCmd)
}
