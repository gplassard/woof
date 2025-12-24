package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentTodosCmd = &cobra.Command{
	Use:   "listincidenttodos",
	Short: "Get a list of an incident's todos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/{incident_id}/relationships/todos")
		fmt.Println("OperationID: ListIncidentTodos")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTodosCmd)
}
