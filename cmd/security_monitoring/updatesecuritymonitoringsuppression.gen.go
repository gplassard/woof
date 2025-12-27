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
	Use:   "updatesecuritymonitoringsuppression [suppression_id]",
	Short: "Update a suppression rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateSecurityMonitoringSuppression(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringSuppressionUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updatesecuritymonitoringsuppression: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityMonitoringSuppressionCmd)
}
