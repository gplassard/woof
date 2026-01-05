package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateMonitorConfigPolicyCmd = &cobra.Command{
	Use:     "update-monitor-config-policy [policy_id]",
	Aliases: []string{"update-config-policy"},
	Short:   "Edit a monitor configuration policy",
	Long: `Edit a monitor configuration policy
Documentation: https://docs.datadoghq.com/api/latest/monitors/#update-monitor-config-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorConfigPolicyResponse
		var err error

		var body datadogV2.MonitorConfigPolicyEditRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err = api.UpdateMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-monitor-config-policy")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-config-policy"))
	},
}

func init() {

	UpdateMonitorConfigPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateMonitorConfigPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateMonitorConfigPolicyCmd)
}
