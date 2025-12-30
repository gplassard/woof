package actions_datastores

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDatastoreItemsCmd = &cobra.Command{
	Use:     "list-datastore-items [datastore_id]",
	Aliases: []string{"list-items"},
	Short:   "List datastore items",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.ListDatastoreItems(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-datastore-items")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {
	Cmd.AddCommand(ListDatastoreItemsCmd)
}
