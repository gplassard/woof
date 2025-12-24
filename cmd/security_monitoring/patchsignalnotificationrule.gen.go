package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var PatchSignalNotificationRuleCmd = &cobra.Command{
	Use:   "patchsignalnotificationrule [id]",
	Short: "Patch a signal-based notification rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.PatchSignalNotificationRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.PatchNotificationRuleParameters{})
		if err != nil {
			log.Fatalf("failed to patchsignalnotificationrule: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(PatchSignalNotificationRuleCmd)
}
