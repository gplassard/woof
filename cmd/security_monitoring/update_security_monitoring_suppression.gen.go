package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:     "update-security-monitoring-suppression [suppression_id] [payload]",
	Aliases: []string{"update-suppression"},
	Short:   "Update a suppression rule",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSuppressionResponse
		var err error

		var body datadogV2.SecurityMonitoringSuppressionUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.UpdateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-security-monitoring-suppression")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityMonitoringSuppressionCmd)
}
