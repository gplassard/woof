package monitors

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateMonitorNotificationRuleCmd = &cobra.Command{
	Use:   "updatemonitornotificationrule [rule_id]",
	Short: "Update a monitor notification rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.UpdateMonitorNotificationRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MonitorNotificationRuleUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updatemonitornotificationrule: %v", err)
		}

		cmdutil.PrintJSON(res, "monitors")
	},
}

func init() {
	Cmd.AddCommand(UpdateMonitorNotificationRuleCmd)
}
