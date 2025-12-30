package monitors

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateMonitorUserTemplateCmd = &cobra.Command{
	Use:     "create-monitor-user-template",
	Aliases: []string{"create-user-template"},
	Short:   "Create a monitor user template",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.CreateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), datadogV2.MonitorUserTemplateCreateRequest{})
		cmdutil.HandleError(err, "failed to create-monitor-user-template")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-user-template"))
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorUserTemplateCmd)
}
