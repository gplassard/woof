package dashboard_lists

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDashboardListItemsCmd = &cobra.Command{
	Use:   "getdashboardlistitems",
	Short: "Get items of a Dashboard List",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards")
		fmt.Println("OperationID: GetDashboardListItems")
	},
}

func init() {
	Cmd.AddCommand(GetDashboardListItemsCmd)
}
