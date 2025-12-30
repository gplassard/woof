package actions_datastores

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var BulkDeleteDatastoreItemsCmd = &cobra.Command{
	Use:     "bulk-delete-datastore-items [datastore_id]",
	Aliases: []string{"bulk-delete-items"},
	Short:   "Bulk delete datastore items",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.BulkDeleteDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], datadogV2.BulkDeleteAppsDatastoreItemsRequest{})
		cmdutil.HandleError(err, "failed to bulk-delete-datastore-items")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {
	Cmd.AddCommand(BulkDeleteDatastoreItemsCmd)
}
