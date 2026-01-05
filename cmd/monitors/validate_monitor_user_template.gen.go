package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateMonitorUserTemplateCmd = &cobra.Command{
	Use:     "validate-monitor-user-template",
	Aliases: []string{"validate-user-template"},
	Short:   "Validate a monitor user template",
	Long: `Validate a monitor user template
Documentation: https://docs.datadoghq.com/api/latest/monitors/#validate-monitor-user-template`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.MonitorUserTemplateCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err = api.ValidateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-monitor-user-template")

	},
}

func init() {

	ValidateMonitorUserTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ValidateMonitorUserTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ValidateMonitorUserTemplateCmd)
}
