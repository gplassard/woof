package monitors

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetMonitorConfigPolicyCmd = &cobra.Command{
	Use:     "get-monitor-config-policy [policy_id]",
	Aliases: []string{"get-config-policy"},
	Short:   "Get a monitor configuration policy",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.GetMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-monitor-config-policy")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-config-policy"))
	},
}

func init() {
	Cmd.AddCommand(GetMonitorConfigPolicyCmd)
}
