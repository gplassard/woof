package monitors

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetMonitorUserTemplateCmd = &cobra.Command{
	Use:     "get-monitor-user-template [template_id]",
	Aliases: []string{"get-user-template"},
	Short:   "Get a monitor user template",
	Long: `Get a monitor user template
Documentation: https://docs.datadoghq.com/api/latest/monitors/#get-monitor-user-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorUserTemplateResponse
		var err error

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-monitor-user-template")

		fmt.Println(cmdutil.FormatJSON(res, "monitor_user_template"))
	},
}

func init() {

	Cmd.AddCommand(GetMonitorUserTemplateCmd)
}
