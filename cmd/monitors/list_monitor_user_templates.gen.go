package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMonitorUserTemplatesCmd = &cobra.Command{
	Use:     "list-monitor-user-templates",
	Aliases: []string{"list-user-templates"},
	Short:   "Get all monitor user templates",
	Long: `Get all monitor user templates
Documentation: https://docs.datadoghq.com/api/latest/monitors/#list-monitor-user-templates`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorUserTemplateListResponse
		var err error

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListMonitorUserTemplates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-monitor-user-templates")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-user-template"))
	},
}

func init() {

	Cmd.AddCommand(ListMonitorUserTemplatesCmd)
}
