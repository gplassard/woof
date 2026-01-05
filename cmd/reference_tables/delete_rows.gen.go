package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteRowsCmd = &cobra.Command{
	Use: "delete-rows [id]",

	Short: "Delete rows",
	Long: `Delete rows
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#delete-rows`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.BatchDeleteRowsRequestArray
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		_, err = api.DeleteRows(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to delete-rows")

	},
}

func init() {

	DeleteRowsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteRowsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteRowsCmd)
}
