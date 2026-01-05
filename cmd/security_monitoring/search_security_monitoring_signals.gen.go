package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:     "search-security-monitoring-signals",
	Aliases: []string{"search-signals"},
	Short:   "Get a list of security signals",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.SearchSecurityMonitoringSignals(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchSecurityMonitoringSignalsOptionalParameters())
		cmdutil.HandleError(err, "failed to search-security-monitoring-signals")

		cmd.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringSignalsCmd)
}
