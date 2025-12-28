package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "update_security_monitoring_suppression [suppression_id]",
	Short: "Update a suppression rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringSuppressionUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_security_monitoring_suppression: %v", err)
		}

		cmdutil.PrintJSON(res, "suppressions")
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityMonitoringSuppressionCmd)
}
