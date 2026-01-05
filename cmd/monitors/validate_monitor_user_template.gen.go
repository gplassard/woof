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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.ValidateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), datadogV2.MonitorUserTemplateCreateRequest{})
		cmdutil.HandleError(err, "failed to validate-monitor-user-template")

	},
}

func init() {
	Cmd.AddCommand(ValidateMonitorUserTemplateCmd)
}
