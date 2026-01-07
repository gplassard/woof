package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringRulesCmd = &cobra.Command{
	Use:     "list-security-monitoring-rules",
	Aliases: []string{"list-rules"},
	Short:   "List rules",
	Long: `List rules
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-security-monitoring-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringListRulesResponse
		var err error

		optionalParams := datadogV2.NewListSecurityMonitoringRulesOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSecurityMonitoringRules(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-security-monitoring-rules")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	ListSecurityMonitoringRulesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListSecurityMonitoringRulesCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	Cmd.AddCommand(ListSecurityMonitoringRulesCmd)
}
