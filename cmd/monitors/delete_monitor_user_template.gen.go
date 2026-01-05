package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteMonitorUserTemplateCmd = &cobra.Command{
	Use:     "delete-monitor-user-template [template_id]",
	Aliases: []string{"delete-user-template"},
	Short:   "Delete a monitor user template",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.DeleteMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-monitor-user-template")

	},
}

func init() {
	Cmd.AddCommand(DeleteMonitorUserTemplateCmd)
}
