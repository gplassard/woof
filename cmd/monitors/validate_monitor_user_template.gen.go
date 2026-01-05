package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ValidateMonitorUserTemplateCmd = &cobra.Command{
	Use:     "validate-monitor-user-template [payload]",
	Aliases: []string{"validate-user-template"},
	Short:   "Validate a monitor user template",
	Long: `Validate a monitor user template
Documentation: https://docs.datadoghq.com/api/latest/monitors/#validate-monitor-user-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.MonitorUserTemplateCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err = api.ValidateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-monitor-user-template")

	},
}

func init() {
	Cmd.AddCommand(ValidateMonitorUserTemplateCmd)
}
