package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpsertRowsCmd = &cobra.Command{
	Use: "upsert-rows [id]",

	Short: "Upsert rows",
	Long: `Upsert rows
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#upsert-rows`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.BatchUpsertRowsRequestArray
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		_, err = api.UpsertRows(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to upsert-rows")

	},
}

func init() {

	UpsertRowsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpsertRowsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpsertRowsCmd)
}
