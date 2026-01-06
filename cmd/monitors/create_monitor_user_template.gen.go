package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateMonitorUserTemplateCmd = &cobra.Command{
	Use:     "create-monitor-user-template",
	Aliases: []string{"create-user-template"},
	Short:   "Create a monitor user template",
	Long: `Create a monitor user template
Documentation: https://docs.datadoghq.com/api/latest/monitors/#create-monitor-user-template`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorUserTemplateCreateResponse
		var err error

		var body datadogV2.MonitorUserTemplateCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-monitor-user-template")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-user-template"))
	},
}

func init() {

	CreateMonitorUserTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateMonitorUserTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateMonitorUserTemplateCmd)
}
