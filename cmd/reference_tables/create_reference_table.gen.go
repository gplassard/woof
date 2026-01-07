package reference_tables

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateReferenceTableCmd = &cobra.Command{
	Use:     "create-reference-table",
	Aliases: []string{"create"},
	Short:   "Create reference table",
	Long: `Create reference table
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#create-reference-table`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TableResultV2
		var err error

		var body datadogV2.CreateTableRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateReferenceTable(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-reference-table")

		fmt.Println(cmdutil.FormatJSON(res, "reference_table"))
	},
}

func init() {

	CreateReferenceTableCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateReferenceTableCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateReferenceTableCmd)
}
