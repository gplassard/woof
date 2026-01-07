package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringSuppressionsCmd = &cobra.Command{
	Use:     "list-security-monitoring-suppressions",
	Aliases: []string{"list-suppressions"},
	Short:   "Get all suppression rules",
	Long: `Get all suppression rules
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-security-monitoring-suppressions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSuppressionsResponse
		var err error

		optionalParams := datadogV2.NewListSecurityMonitoringSuppressionsOptionalParameters()

		if cmd.Flags().Changed("query") {
			val, _ := cmd.Flags().GetString("query")
			optionalParams.WithQuery(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSecurityMonitoringSuppressions(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-security-monitoring-suppressions")

		cmd.Println(cmdutil.FormatJSON(res, "suppressions"))
	},
}

func init() {

	ListSecurityMonitoringSuppressionsCmd.Flags().String("query", "", "Query string.")

	Cmd.AddCommand(ListSecurityMonitoringSuppressionsCmd)
}
