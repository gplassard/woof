package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var EditSecurityMonitoringSignalStateCmd = &cobra.Command{
	Use:     "edit-security-monitoring-signal-state [signal_id]",
	Aliases: []string{"edit-signal-state"},
	Short:   "Change the triage state of a security signal",
	Long: `Change the triage state of a security signal
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#edit-security-monitoring-signal-state`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalTriageUpdateResponse
		var err error

		var body datadogV2.SecurityMonitoringSignalStateUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.EditSecurityMonitoringSignalState(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to edit-security-monitoring-signal-state")

		fmt.Println(cmdutil.FormatJSON(res, "signal_metadata"))
	},
}

func init() {

	EditSecurityMonitoringSignalStateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	EditSecurityMonitoringSignalStateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(EditSecurityMonitoringSignalStateCmd)
}
