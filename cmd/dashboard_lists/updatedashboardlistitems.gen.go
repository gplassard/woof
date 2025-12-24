package dashboard_lists

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDashboardListItemsCmd = &cobra.Command{
	Use:   "updatedashboardlistitems",
	Short: "Update items of a dashboard list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards")
		fmt.Println("OperationID: UpdateDashboardListItems")
	},
}

func init() {
	Cmd.AddCommand(UpdateDashboardListItemsCmd)
}
