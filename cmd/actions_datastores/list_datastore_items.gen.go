package actions_datastores

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDatastoreItemsCmd = &cobra.Command{
	Use:     "list-datastore-items [datastore_id]",
	Aliases: []string{"list-items"},
	Short:   "List datastore items",
	Long: `List datastore items
Documentation: https://docs.datadoghq.com/api/latest/actions-datastores/#list-datastore-items`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ItemApiPayloadArray
		var err error

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err = api.ListDatastoreItems(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-datastore-items")

		cmd.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {
	Cmd.AddCommand(ListDatastoreItemsCmd)
}
