package dashboard_lists

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDashboardListItemsCmd = &cobra.Command{
	Use:   "createdashboardlistitems",
	Short: "Add Items to a Dashboard List",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards")
		fmt.Println("OperationID: CreateDashboardListItems")
	},
}

func init() {
	Cmd.AddCommand(CreateDashboardListItemsCmd)
}
