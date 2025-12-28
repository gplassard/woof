package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "get-security-monitoring-suppression [suppression_id]",
	Aliases: []string{ "get-suppression", },
	Short: "Get a suppression rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-security-monitoring-suppression")

		cmdutil.PrintJSON(res, "suppressions")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringSuppressionCmd)
}
