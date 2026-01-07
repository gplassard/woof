package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecurityMonitoringHistsignalsByJobIdCmd = &cobra.Command{
	Use:     "get-security-monitoring-histsignals-by-job-id [job_id]",
	Aliases: []string{"get-histsignals-by-job-id"},
	Short:   "Get a job's hist signals",
	Long: `Get a job's hist signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-security-monitoring-histsignals-by-job-id`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		optionalParams := datadogV2.NewGetSecurityMonitoringHistsignalsByJobIdOptionalParameters()

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
		res, _, err = api.GetSecurityMonitoringHistsignalsByJobId(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-security-monitoring-histsignals-by-job-id")

		cmd.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {

	GetSecurityMonitoringHistsignalsByJobIdCmd.Flags().String("filter-query", "", "The search query for security signals.")

	GetSecurityMonitoringHistsignalsByJobIdCmd.Flags().String("filter-from", "", "The minimum timestamp for requested security signals.")

	GetSecurityMonitoringHistsignalsByJobIdCmd.Flags().String("filter-to", "", "The maximum timestamp for requested security signals.")

	GetSecurityMonitoringHistsignalsByJobIdCmd.Flags().String("sort", "", "The order of the security signals in results.")

	GetSecurityMonitoringHistsignalsByJobIdCmd.Flags().String("page-cursor", "", "A list of results using the cursor provided in the previous query.")

	GetSecurityMonitoringHistsignalsByJobIdCmd.Flags().Int64("page-limit", 0, "The maximum number of security signals in the response.")

	Cmd.AddCommand(GetSecurityMonitoringHistsignalsByJobIdCmd)
}
