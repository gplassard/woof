package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteMonitorConfigPolicyCmd = &cobra.Command{
	Use:     "delete-monitor-config-policy [policy_id]",
	Aliases: []string{"delete-config-policy"},
	Short:   "Delete a monitor configuration policy",
	Long: `Delete a monitor configuration policy
Documentation: https://docs.datadoghq.com/api/latest/monitors/#delete-monitor-config-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-monitor-config-policy")

	},
}

func init() {

	Cmd.AddCommand(DeleteMonitorConfigPolicyCmd)
}
