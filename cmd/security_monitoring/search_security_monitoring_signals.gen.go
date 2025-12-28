package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:   "search-security-monitoring-signals",
	Short: "Get a list of security signals",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.SearchSecurityMonitoringSignals(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchSecurityMonitoringSignalsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to search-security-monitoring-signals: %v", err)
		}

		cmdutil.PrintJSON(res, "signal")
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringSignalsCmd)
}
