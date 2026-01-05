package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var BulkDeleteDatastoreItemsCmd = &cobra.Command{
	Use:     "bulk-delete-datastore-items [datastore_id] [payload]",
	Aliases: []string{"bulk-delete-items"},
	Short:   "Bulk delete datastore items",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.BulkDeleteAppsDatastoreItemsRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.BulkDeleteDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to bulk-delete-datastore-items")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {
	Cmd.AddCommand(BulkDeleteDatastoreItemsCmd)
}
