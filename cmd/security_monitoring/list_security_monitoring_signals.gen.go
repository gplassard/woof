package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:     "list-security-monitoring-signals",
	Aliases: []string{"list-signals"},
	Short:   "Get a quick list of security signals",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityMonitoringSignals(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-signals")

		cmd.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringSignalsCmd)
}
