package actions_datastores

import (
	"fmt"
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

		optionalParams := datadogV2.NewListDatastoreItemsOptionalParameters()

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		if cmd.Flags().Changed("item-key") {
			val, _ := cmd.Flags().GetString("item-key")
			optionalParams.WithItemKey(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to list-datastore-items")

		fmt.Println(cmdutil.FormatJSON(res, "items"))
	},
}

func init() {

	ListDatastoreItemsCmd.Flags().String("filter", "", "Optional query filter to search items using the [logs search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax/).")

	ListDatastoreItemsCmd.Flags().String("item-key", "", "Optional primary key value to retrieve a specific item. Cannot be used together with the filter parameter.")

	ListDatastoreItemsCmd.Flags().Int64("page-limit", 0, "Optional field to limit the number of items to return per page for pagination. Up to 100 items can be returned per page.")

	ListDatastoreItemsCmd.Flags().Int64("page-offset", 0, "Optional field to offset the number of items to skip from the beginning of the result set for pagination.")

	ListDatastoreItemsCmd.Flags().String("sort", "", "Optional field to sort results by. Prefix with '-' for descending order (e.g., '-created_at').")

	Cmd.AddCommand(ListDatastoreItemsCmd)
}
