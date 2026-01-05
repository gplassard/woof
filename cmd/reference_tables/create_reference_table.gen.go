package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateReferenceTableCmd = &cobra.Command{
	Use:     "create-reference-table [payload]",
	Aliases: []string{"create"},
	Short:   "Create reference table",
	Long: `Create reference table
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#create-reference-table`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TableResultV2
		var err error

		var body datadogV2.CreateTableRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err = api.CreateReferenceTable(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-reference-table")

		cmd.Println(cmdutil.FormatJSON(res, "reference_table"))
	},
}

func init() {
	Cmd.AddCommand(CreateReferenceTableCmd)
}
