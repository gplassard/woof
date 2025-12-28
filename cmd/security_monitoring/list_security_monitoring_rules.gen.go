package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListSecurityMonitoringRulesCmd = &cobra.Command{
	Use:   "list-security-monitoring-rules",
	Aliases: []string{ "list-rules", },
	Short: "List rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityMonitoringRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-security-monitoring-rules: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringRulesCmd)
}
