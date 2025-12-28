package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var EditSecurityMonitoringSignalStateCmd = &cobra.Command{
	Use:   "edit-security-monitoring-signal-state [signal_id]",
	Aliases: []string{ "edit-signal-state", },
	Short: "Change the triage state of a security signal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.EditSecurityMonitoringSignalState(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringSignalStateUpdateRequest{})
		cmdutil.HandleError(err, "failed to edit-security-monitoring-signal-state")

		cmdutil.PrintJSON(res, "signal_metadata")
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalStateCmd)
}
