package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateMonitorUserTemplateCmd = &cobra.Command{
	Use:     "update-monitor-user-template [template_id]",
	Aliases: []string{"update-user-template"},
	Short:   "Update a monitor user template to a new version",
	Long: `Update a monitor user template to a new version
Documentation: https://docs.datadoghq.com/api/latest/monitors/#update-monitor-user-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorUserTemplateResponse
		var err error

		var body datadogV2.MonitorUserTemplateUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err = api.UpdateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-monitor-user-template")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-user-template"))
	},
}

func init() {

	UpdateMonitorUserTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateMonitorUserTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateMonitorUserTemplateCmd)
}
