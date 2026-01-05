package dashboard_lists

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var DeleteDashboardListItemsCmd = &cobra.Command{
	Use:     "delete-dashboard-list-items [dashboard_list_id]",
	Aliases: []string{"delete-items"},
	Short:   "Delete items from a dashboard list",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		res, _, err := api.DeleteDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.DashboardListDeleteItemsRequest{})
		cmdutil.HandleError(err, "failed to delete-dashboard-list-items")

		cmd.Println(cmdutil.FormatJSON(res, "dashboard_lists"))
	},
}

func init() {
	Cmd.AddCommand(DeleteDashboardListItemsCmd)
}
