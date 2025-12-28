package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListSecurityMonitoringSuppressionsCmd = &cobra.Command{
	Use:   "list-security-monitoring-suppressions",
	Aliases: []string{ "list-suppressions", },
	Short: "Get all suppression rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityMonitoringSuppressions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-security-monitoring-suppressions: %v", err)
		}

		cmdutil.PrintJSON(res, "suppressions")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringSuppressionsCmd)
}
