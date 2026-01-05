package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:     "list-security-monitoring-histsignals",
	Aliases: []string{"list-histsignals"},
	Short:   "List hist signals",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityMonitoringHistsignals(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-histsignals")

		cmd.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringHistsignalsCmd)
}
