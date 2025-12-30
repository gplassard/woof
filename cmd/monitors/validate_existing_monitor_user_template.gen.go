package monitors

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateExistingMonitorUserTemplateCmd = &cobra.Command{
	Use:     "validate-existing-monitor-user-template [template_id]",
	Aliases: []string{"validate-existing-user-template"},
	Short:   "Validate an existing monitor user template",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.ValidateExistingMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MonitorUserTemplateUpdateRequest{})
		cmdutil.HandleError(err, "failed to validate-existing-monitor-user-template")

	},
}

func init() {
	Cmd.AddCommand(ValidateExistingMonitorUserTemplateCmd)
}
