package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateExistingMonitorUserTemplateCmd = &cobra.Command{
	Use:     "validate-existing-monitor-user-template [template_id]",
	Aliases: []string{"validate-existing-user-template"},
	Short:   "Validate an existing monitor user template",
	Long: `Validate an existing monitor user template
Documentation: https://docs.datadoghq.com/api/latest/monitors/#validate-existing-monitor-user-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.MonitorUserTemplateUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err = api.ValidateExistingMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to validate-existing-monitor-user-template")

	},
}

func init() {

	ValidateExistingMonitorUserTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ValidateExistingMonitorUserTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ValidateExistingMonitorUserTemplateCmd)
}
