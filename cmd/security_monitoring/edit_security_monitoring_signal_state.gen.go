package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var EditSecurityMonitoringSignalStateCmd = &cobra.Command{
	Use:     "edit-security-monitoring-signal-state [signal_id] [payload]",
	Aliases: []string{"edit-signal-state"},
	Short:   "Change the triage state of a security signal",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SecurityMonitoringSignalStateUpdateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.EditSecurityMonitoringSignalState(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to edit-security-monitoring-signal-state")

		cmd.Println(cmdutil.FormatJSON(res, "signal_metadata"))
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalStateCmd)
}
