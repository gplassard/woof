package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:   "search-security-monitoring-histsignals",
	Aliases: []string{ "search-histsignals", },
	Short: "Search hist signals",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.SearchSecurityMonitoringHistsignals(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchSecurityMonitoringHistsignalsOptionalParameters())
		cmdutil.HandleError(err, "failed to search-security-monitoring-histsignals")

		cmdutil.PrintJSON(res, "signal")
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringHistsignalsCmd)
}
