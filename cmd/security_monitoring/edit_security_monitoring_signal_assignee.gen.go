package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var EditSecurityMonitoringSignalAssigneeCmd = &cobra.Command{
	Use:     "edit-security-monitoring-signal-assignee [signal_id] [payload]",
	Aliases: []string{"edit-signal-assignee"},
	Short:   "Modify the triage assignee of a security signal",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SecurityMonitoringSignalAssigneeUpdateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.EditSecurityMonitoringSignalAssignee(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to edit-security-monitoring-signal-assignee")

		cmd.Println(cmdutil.FormatJSON(res, "signal_metadata"))
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalAssigneeCmd)
}
