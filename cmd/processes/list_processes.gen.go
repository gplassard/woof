package processes

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListProcessesCmd = &cobra.Command{
	Use:     "list-processes",
	Aliases: []string{"list"},
	Short:   "Get all processes",
	Long: `Get all processes
Documentation: https://docs.datadoghq.com/api/latest/processes/#list-processes`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProcessSummariesResponse
		var err error

		optionalParams := datadogV2.NewListProcessesOptionalParameters()

		if cmd.Flags().Changed("search") {
			val, _ := cmd.Flags().GetString("search")
			optionalParams.WithSearch(val)
		}

		if cmd.Flags().Changed("tags") {
			val, _ := cmd.Flags().GetString("tags")
			optionalParams.WithTags(val)
		}

		if cmd.Flags().Changed("from") {
			val, _ := cmd.Flags().GetInt64("from")
			optionalParams.WithFrom(val)
		}

		if cmd.Flags().Changed("to") {
			val, _ := cmd.Flags().GetInt64("to")
			optionalParams.WithTo(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("page-cursor") {
			val, _ := cmd.Flags().GetString("page-cursor")
			optionalParams.WithPageCursor(val)
		}

		api := datadogV2.NewProcessesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListProcesses(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-processes")

		fmt.Println(cmdutil.FormatJSON(res, "process"))
	},
}

func init() {

	ListProcessesCmd.Flags().String("search", "", "String to search processes by.")

	ListProcessesCmd.Flags().String("tags", "", "Comma-separated list of tags to filter processes by.")

	ListProcessesCmd.Flags().Int64("from", 0, "Unix timestamp (number of seconds since epoch) of the start of the query window. If not provided, the start of the query window will be 15 minutes before the 'to' timestamp. If neither 'from' nor 'to' are provided, the query window will be '[now - 15m, now]'.")

	ListProcessesCmd.Flags().Int64("to", 0, "Unix timestamp (number of seconds since epoch) of the end of the query window. If not provided, the end of the query window will be 15 minutes after the 'from' timestamp. If neither 'from' nor 'to' are provided, the query window will be '[now - 15m, now]'.")

	ListProcessesCmd.Flags().Int64("page-limit", 0, "Maximum number of results returned.")

	ListProcessesCmd.Flags().String("page-cursor", "", "String to query the next page of results. This key is provided with each valid response from the API in 'meta.page.after'.")

	Cmd.AddCommand(ListProcessesCmd)
}
