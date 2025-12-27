package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListMultipleRulesetsCmd = &cobra.Command{
	Use:   "listmultiplerulesets",
	Short: "Ruleset get multiple",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListMultipleRulesets(client.NewContext(apiKey, appKey, site), datadogV2.GetMultipleRulesetsRequest{})
		if err != nil {
			log.Fatalf("failed to listmultiplerulesets: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ListMultipleRulesetsCmd)
}
