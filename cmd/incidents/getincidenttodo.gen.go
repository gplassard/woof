package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentTodoCmd = &cobra.Command{
	Use:   "getincidenttodo",
	Short: "Get incident todo details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/{incident_id}/relationships/todos/{todo_id}")
		fmt.Println("OperationID: GetIncidentTodo")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTodoCmd)
}
