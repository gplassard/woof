package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTablesCmd = &cobra.Command{
	Use: "list-tables",

	Short: "List tables",
	Long: `List tables
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#list-tables`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TableResultV2Array
		var err error

		optionalParams := datadogV2.NewListTablesOptionalParameters()

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("filter-status") {
			val, _ := cmd.Flags().GetString("filter-status")
			optionalParams.WithFilterStatus(val)
		}

		if cmd.Flags().Changed("filter-table-name-exact") {
			val, _ := cmd.Flags().GetString("filter-table-name-exact")
			optionalParams.WithFilterTableNameExact(val)
		}

		if cmd.Flags().Changed("filter-table-name-contains") {
			val, _ := cmd.Flags().GetString("filter-table-name-contains")
			optionalParams.WithFilterTableNameContains(val)
		}

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTables(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-tables")

		cmd.Println(cmdutil.FormatJSON(res, "reference_table"))
	},
}

func init() {

	ListTablesCmd.Flags().Int64("page-limit", 0, "Number of tables to return.")

	ListTablesCmd.Flags().Int64("page-offset", 0, "Number of tables to skip for pagination.")

	ListTablesCmd.Flags().String("sort", "", "Sort field and direction for the list of reference tables. Use field name for ascending, prefix with \"-\" for descending.")

	ListTablesCmd.Flags().String("filter-status", "", "Filter by table status.")

	ListTablesCmd.Flags().String("filter-table-name-exact", "", "Filter by exact table name match.")

	ListTablesCmd.Flags().String("filter-table-name-contains", "", "Filter by table name containing substring.")

	Cmd.AddCommand(ListTablesCmd)
}
