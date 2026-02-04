package dashboard_lists

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var GetDashboardListItemsCmd = &cobra.Command{
	Use:     "get-dashboard-list-items [dashboard_list_id]",
	Aliases: []string{"get-items"},
	Short:   "Get items of a Dashboard List",
	Long: `Get items of a Dashboard List
Documentation: https://docs.datadoghq.com/api/latest/dashboard-lists/#get-dashboard-list-items`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DashboardListItems
		var err error

		api := datadogV2.NewDashboardListsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDashboardListItems(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-dashboard-list-items")

		fmt.Println(cmdutil.FormatJSON(res, "dashboard_list_item"))
	},
}

func init() {

	Cmd.AddCommand(GetDashboardListItemsCmd)
}
