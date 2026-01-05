package monitors

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMonitorConfigPoliciesCmd = &cobra.Command{
	Use:     "list-monitor-config-policies",
	Aliases: []string{"list-config-policies"},
	Short:   "Get all monitor configuration policies",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.ListMonitorConfigPolicies(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-monitor-config-policies")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-config-policy"))
	},
}

func init() {
	Cmd.AddCommand(ListMonitorConfigPoliciesCmd)
}
