package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringSuppressionsCmd = &cobra.Command{
	Use:     "list-security-monitoring-suppressions",
	Aliases: []string{"list-suppressions"},
	Short:   "Get all suppression rules",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityMonitoringSuppressions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-suppressions")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringSuppressionsCmd)
}
