package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSignalNotificationRuleCmd = &cobra.Command{
	Use:   "createsignalnotificationrule",
	Short: "Create a new signal-based notification rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateSignalNotificationRule(client.NewContext(apiKey, appKey, site), datadogV2.CreateNotificationRuleParameters{})
		if err != nil {
			log.Fatalf("failed to createsignalnotificationrule: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(CreateSignalNotificationRuleCmd)
}
