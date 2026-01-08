package monitors

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMonitorConfigPoliciesCmd = &cobra.Command{
	Use:     "list-monitor-config-policies",
	Aliases: []string{"list-config-policies"},
	Short:   "Get all monitor configuration policies",
	Long: `Get all monitor configuration policies
Documentation: https://docs.datadoghq.com/api/latest/monitors/#list-monitor-config-policies`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorConfigPolicyListResponse
		var err error

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListMonitorConfigPolicies(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-monitor-config-policies")

		fmt.Println(cmdutil.FormatJSON(res, "monitor_config_policie"))
	},
}

func init() {

	Cmd.AddCommand(ListMonitorConfigPoliciesCmd)
}
