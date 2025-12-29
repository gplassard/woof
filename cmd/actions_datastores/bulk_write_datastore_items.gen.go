package actions_datastores

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var BulkWriteDatastoreItemsCmd = &cobra.Command{
	Use:   "bulk-write-datastore-items [datastore_id]",
	Aliases: []string{ "bulk-write-items", },
	Short: "Bulk write datastore items",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.BulkWriteDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], datadogV2.BulkPutAppsDatastoreItemsRequest{})
		cmdutil.HandleError(err, "failed to bulk-write-datastore-items")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {
	Cmd.AddCommand(BulkWriteDatastoreItemsCmd)
}
