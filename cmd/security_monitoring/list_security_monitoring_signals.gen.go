package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:     "list-security-monitoring-signals",
	Aliases: []string{"list-signals"},
	Short:   "Get a quick list of security signals",
	Long: `Get a quick list of security signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-security-monitoring-signals`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		optionalParams := datadogV2.NewListSecurityMonitoringSignalsOptionalParameters()

		if cmd.Flags().Changed("filter-query") {
			val, _ := cmd.Flags().GetString("filter-query")
			optionalParams.WithFilterQuery(val)
		}

		if cmd.Flags().Changed("filter-from") {
			val, _ := cmd.Flags().GetString("filter-from")
			optionalParams.WithFilterFrom(val)
		}

		if cmd.Flags().Changed("filter-to") {
			val, _ := cmd.Flags().GetString("filter-to")
			optionalParams.WithFilterTo(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("page-cursor") {
			val, _ := cmd.Flags().GetString("page-cursor")
			optionalParams.WithPageCursor(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSecurityMonitoringSignals(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-security-monitoring-signals")

		fmt.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {

	ListSecurityMonitoringSignalsCmd.Flags().String("filter-query", "", "The search query for security signals.")

	ListSecurityMonitoringSignalsCmd.Flags().String("filter-from", "", "The minimum timestamp for requested security signals.")

	ListSecurityMonitoringSignalsCmd.Flags().String("filter-to", "", "The maximum timestamp for requested security signals.")

	ListSecurityMonitoringSignalsCmd.Flags().String("sort", "", "The order of the security signals in results.")

	ListSecurityMonitoringSignalsCmd.Flags().String("page-cursor", "", "A list of results using the cursor provided in the previous query.")

	ListSecurityMonitoringSignalsCmd.Flags().Int64("page-limit", 0, "The maximum number of security signals in the response.")

	Cmd.AddCommand(ListSecurityMonitoringSignalsCmd)
}
