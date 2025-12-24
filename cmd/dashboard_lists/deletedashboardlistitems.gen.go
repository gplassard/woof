package dashboard_lists

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDashboardListItemsCmd = &cobra.Command{
	Use:   "deletedashboardlistitems",
	Short: "Delete items from a dashboard list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards")
		fmt.Println("OperationID: DeleteDashboardListItems")
	},
}

func init() {
	Cmd.AddCommand(DeleteDashboardListItemsCmd)
}
