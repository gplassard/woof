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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		_, err := api.DeleteRows(client.NewContext(apiKey, appKey, site), args[0], datadogV2.BatchDeleteRowsRequestArray{})
		cmdutil.HandleError(err, "failed to delete-rows")

	},
}

func init() {
	Cmd.AddCommand(DeleteRowsCmd)
}
