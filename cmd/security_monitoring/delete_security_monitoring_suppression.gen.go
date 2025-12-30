package security_monitoring

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:     "delete-security-monitoring-suppression [suppression_id]",
	Aliases: []string{"delete-suppression"},
	Short:   "Delete a suppression rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DeleteSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-security-monitoring-suppression")

	},
}

func init() {
	Cmd.AddCommand(DeleteSecurityMonitoringSuppressionCmd)
}
