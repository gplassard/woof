package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var QueryEventFilteredUsersCmd = &cobra.Command{
	Use:   "queryeventfilteredusers",
	Short: "Query event filtered users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/product-analytics/users/event_filtered_query")
		fmt.Println("OperationID: QueryEventFilteredUsers")
	},
}

func init() {
	Cmd.AddCommand(QueryEventFilteredUsersCmd)
}
