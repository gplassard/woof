package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateMonitorUserTemplateCmd = &cobra.Command{
	Use:     "create-monitor-user-template [payload]",
	Aliases: []string{"create-user-template"},
	Short:   "Create a monitor user template",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.MonitorUserTemplateCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.CreateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-monitor-user-template")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-user-template"))
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorUserTemplateCmd)
}
