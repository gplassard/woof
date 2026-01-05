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

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err = api.ListTables(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-tables")

		cmd.Println(cmdutil.FormatJSON(res, "reference_table"))
	},
}

func init() {
	Cmd.AddCommand(ListTablesCmd)
}
