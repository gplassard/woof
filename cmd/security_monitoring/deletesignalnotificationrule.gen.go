package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteSignalNotificationRuleCmd = &cobra.Command{
	Use:   "deletesignalnotificationrule [id]",
	Short: "Delete a signal-based notification rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DeleteSignalNotificationRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletesignalnotificationrule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteSignalNotificationRuleCmd)
}
