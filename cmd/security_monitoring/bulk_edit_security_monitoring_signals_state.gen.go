package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var BulkEditSecurityMonitoringSignalsStateCmd = &cobra.Command{
	Use:     "bulk-edit-security-monitoring-signals-state",
	Aliases: []string{"bulk-edit-signals-state"},
	Short:   "Bulk update triage state of security signals",
	Long: `Bulk update triage state of security signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#bulk-edit-security-monitoring-signals-state`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsBulkTriageUpdateResponse
		var err error

		var body datadogV2.SecurityMonitoringSignalsBulkStateUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.BulkEditSecurityMonitoringSignalsState(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to bulk-edit-security-monitoring-signals-state")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	BulkEditSecurityMonitoringSignalsStateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	BulkEditSecurityMonitoringSignalsStateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(BulkEditSecurityMonitoringSignalsStateCmd)
}
