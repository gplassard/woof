package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetMonitorNotificationRulesCmd = &cobra.Command{
	Use:     "get-monitor-notification-rules",
	Aliases: []string{"get-notification-rules"},
	Short:   "Get all monitor notification rules",
	Long: `Get all monitor notification rules
Documentation: https://docs.datadoghq.com/api/latest/monitors/#get-monitor-notification-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorNotificationRuleListResponse
		var err error

		optionalParams := datadogV2.NewGetMonitorNotificationRulesOptionalParameters()

		if cmd.Flags().Changed("page") {
			val, _ := cmd.Flags().GetInt64("page")
			optionalParams.WithPage(val)
		}

		if cmd.Flags().Changed("per-page") {
			val, _ := cmd.Flags().GetInt64("per-page")
			optionalParams.WithPerPage(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("filters") {
			val, _ := cmd.Flags().GetString("filters")
			optionalParams.WithFilters(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMonitorNotificationRules(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to get-monitor-notification-rules")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {

	GetMonitorNotificationRulesCmd.Flags().Int64("page", 0, "The page to start paginating from. If 'page' is not specified, the argument defaults to the first page.")

	GetMonitorNotificationRulesCmd.Flags().Int64("per-page", 0, "The number of rules to return per page. If 'per_page' is not specified, the argument defaults to 100.")

	GetMonitorNotificationRulesCmd.Flags().String("sort", "", "String for sort order, composed of field and sort order separated by a colon, for example 'name:asc'. Supported sort directions: 'asc', 'desc'. Supported fields: 'name', 'created_at'.")

	GetMonitorNotificationRulesCmd.Flags().String("filters", "", "JSON-encoded filter object. Supported keys: * 'text': Free-text query matched against rule name, tags, and recipients. * 'tags': Array of strings. Return rules that have any of these tags. * 'recipients': Array of strings. Return rules that have any of these recipients.")

	GetMonitorNotificationRulesCmd.Flags().String("include", "", "Comma-separated list of resource paths for related resources to include in the response. Supported resource path is 'created_by'.")

	Cmd.AddCommand(GetMonitorNotificationRulesCmd)
}
