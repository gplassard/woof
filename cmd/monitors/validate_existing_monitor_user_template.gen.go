package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ValidateExistingMonitorUserTemplateCmd = &cobra.Command{
	Use:     "validate-existing-monitor-user-template [template_id] [payload]",
	Aliases: []string{"validate-existing-user-template"},
	Short:   "Validate an existing monitor user template",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.MonitorUserTemplateUpdateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.ValidateExistingMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to validate-existing-monitor-user-template")

	},
}

func init() {
	Cmd.AddCommand(ValidateExistingMonitorUserTemplateCmd)
}
