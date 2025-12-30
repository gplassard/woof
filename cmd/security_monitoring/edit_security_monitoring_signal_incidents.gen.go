package security_monitoring

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var EditSecurityMonitoringSignalIncidentsCmd = &cobra.Command{
	Use:     "edit-security-monitoring-signal-incidents [signal_id]",
	Aliases: []string{"edit-signal-incidents"},
	Short:   "Change the related incidents of a security signal",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.EditSecurityMonitoringSignalIncidents(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringSignalIncidentsUpdateRequest{})
		cmdutil.HandleError(err, "failed to edit-security-monitoring-signal-incidents")

		cmd.Println(cmdutil.FormatJSON(res, "signal_metadata"))
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalIncidentsCmd)
}
