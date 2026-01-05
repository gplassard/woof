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
	Long: `Create a monitor configuration policy
Documentation: https://docs.datadoghq.com/api/latest/monitors/#create-monitor-config-policy`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorConfigPolicyResponse
		var err error

		var body datadogV2.MonitorConfigPolicyCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err = api.CreateMonitorConfigPolicy(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-monitor-config-policy")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-config-policy"))
	},
}

func init() {

	CreateMonitorConfigPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateMonitorConfigPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateMonitorConfigPolicyCmd)
}
