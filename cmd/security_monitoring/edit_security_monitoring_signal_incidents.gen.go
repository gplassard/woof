package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var EditSecurityMonitoringSignalIncidentsCmd = &cobra.Command{
	Use:     "edit-security-monitoring-signal-incidents [signal_id]",
	Aliases: []string{"edit-signal-incidents"},
	Short:   "Change the related incidents of a security signal",
	Long: `Change the related incidents of a security signal
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#edit-security-monitoring-signal-incidents`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalTriageUpdateResponse
		var err error

		var body datadogV2.SecurityMonitoringSignalIncidentsUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.EditSecurityMonitoringSignalIncidents(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to edit-security-monitoring-signal-incidents")

		cmd.Println(cmdutil.FormatJSON(res, "signal_metadata"))
	},
}

func init() {

	EditSecurityMonitoringSignalIncidentsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	EditSecurityMonitoringSignalIncidentsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(EditSecurityMonitoringSignalIncidentsCmd)
}
