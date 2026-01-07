package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListThreatHuntingJobsCmd = &cobra.Command{
	Use: "list-threat-hunting-jobs",

	Short: "List threat hunting jobs",
	Long: `List threat hunting jobs
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-threat-hunting-jobs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListThreatHuntingJobsResponse
		var err error

		optionalParams := datadogV2.NewListThreatHuntingJobsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("filter-query") {
			val, _ := cmd.Flags().GetString("filter-query")
			optionalParams.WithFilterQuery(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListThreatHuntingJobs(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-threat-hunting-jobs")

		cmd.Println(cmdutil.FormatJSON(res, "historicalDetectionsJob"))
	},
}

func init() {

	ListThreatHuntingJobsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListThreatHuntingJobsCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListThreatHuntingJobsCmd.Flags().String("sort", "", "The order of the jobs in results.")

	ListThreatHuntingJobsCmd.Flags().String("filter-query", "", "Query used to filter items from the fetched list.")

	Cmd.AddCommand(ListThreatHuntingJobsCmd)
}
