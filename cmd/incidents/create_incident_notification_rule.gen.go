package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentNotificationRuleCmd = &cobra.Command{
	Use:   "create-incident-notification-rule",
	Aliases: []string{ "create-notification-rule", },
	Short: "Create an incident notification rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentNotificationRule(client.NewContext(apiKey, appKey, site), datadogV2.CreateIncidentNotificationRuleRequest{})
		if err != nil {
			log.Fatalf("failed to create-incident-notification-rule: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_notification_rules")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationRuleCmd)
}
