package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentTodoCmd = &cobra.Command{
	Use:   "updateincidenttodo",
	Short: "Update an incident todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/incidents/{incident_id}/relationships/todos/{todo_id}")
		fmt.Println("OperationID: UpdateIncidentTodo")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTodoCmd)
}
