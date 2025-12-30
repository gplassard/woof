package dashboard_lists

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var CreateDashboardListItemsCmd = &cobra.Command{
	Use:     "create-dashboard-list-items [dashboard_list_id]",
	Aliases: []string{"create-items"},
	Short:   "Add Items to a Dashboard List",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		res, _, err := api.CreateDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), datadogV2.DashboardListAddItemsRequest{})
		cmdutil.HandleError(err, "failed to create-dashboard-list-items")

		cmd.Println(cmdutil.FormatJSON(res, "dashboard_lists"))
	},
}

func init() {
	Cmd.AddCommand(CreateDashboardListItemsCmd)
}
