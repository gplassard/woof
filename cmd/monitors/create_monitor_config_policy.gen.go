package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateMonitorConfigPolicyCmd = &cobra.Command{
	Use:     "create-monitor-config-policy",
	Aliases: []string{"create-config-policy"},
	Short:   "Create a monitor configuration policy",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.CreateMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), datadogV2.MonitorConfigPolicyCreateRequest{})
		cmdutil.HandleError(err, "failed to create-monitor-config-policy")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-config-policy"))
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorConfigPolicyCmd)
}
