package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var EditSecurityMonitoringSignalAssigneeCmd = &cobra.Command{
	Use:   "edit-security-monitoring-signal-assignee [signal_id]",
	Aliases: []string{ "edit-signal-assignee", },
	Short: "Modify the triage assignee of a security signal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.EditSecurityMonitoringSignalAssignee(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringSignalAssigneeUpdateRequest{})
		cmdutil.HandleError(err, "failed to edit-security-monitoring-signal-assignee")

		cmd.Println(cmdutil.FormatJSON(res, "signal_metadata"))
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalAssigneeCmd)
}
