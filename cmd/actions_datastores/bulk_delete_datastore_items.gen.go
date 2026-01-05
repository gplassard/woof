package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var BulkDeleteDatastoreItemsCmd = &cobra.Command{
	Use:     "bulk-delete-datastore-items [datastore_id]",
	Aliases: []string{"bulk-delete-items"},
	Short:   "Bulk delete datastore items",
	Long: `Bulk delete datastore items
Documentation: https://docs.datadoghq.com/api/latest/actions-datastores/#bulk-delete-datastore-items`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeleteAppsDatastoreItemResponseArray
		var err error

		var body datadogV2.BulkDeleteAppsDatastoreItemsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err = api.BulkDeleteDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to bulk-delete-datastore-items")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {

	BulkDeleteDatastoreItemsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	BulkDeleteDatastoreItemsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(BulkDeleteDatastoreItemsCmd)
}
