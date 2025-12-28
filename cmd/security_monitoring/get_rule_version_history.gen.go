package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRuleVersionHistoryCmd = &cobra.Command{
	Use:   "get-rule-version-history [rule_id]",
	Short: "Get a rule's version history",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetRuleVersionHistory(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-rule-version-history: %v", err)
		}

		cmdutil.PrintJSON(res, "GetRuleVersionHistoryResponse")
	},
}

func init() {
	Cmd.AddCommand(GetRuleVersionHistoryCmd)
}
