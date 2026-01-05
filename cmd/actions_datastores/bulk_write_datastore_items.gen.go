package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var BulkWriteDatastoreItemsCmd = &cobra.Command{
	Use:     "bulk-write-datastore-items [datastore_id] [payload]",
	Aliases: []string{"bulk-write-items"},
	Short:   "Bulk write datastore items",
	Long: `Bulk write datastore items
Documentation: https://docs.datadoghq.com/api/latest/actions-datastores/#bulk-write-datastore-items`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PutAppsDatastoreItemResponseArray
		var err error

		var body datadogV2.BulkPutAppsDatastoreItemsRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err = api.BulkWriteDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to bulk-write-datastore-items")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {
	Cmd.AddCommand(BulkWriteDatastoreItemsCmd)
}
