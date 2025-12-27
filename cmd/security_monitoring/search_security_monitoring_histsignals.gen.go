package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:   "search_security_monitoring_histsignals",
	Short: "Search hist signals",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.SearchSecurityMonitoringHistsignals(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchSecurityMonitoringHistsignalsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to search_security_monitoring_histsignals: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringHistsignalsCmd)
}
