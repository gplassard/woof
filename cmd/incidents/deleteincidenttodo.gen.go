package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentTodoCmd = &cobra.Command{
	Use:   "deleteincidenttodo",
	Short: "Delete an incident todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/{incident_id}/relationships/todos/{todo_id}")
		fmt.Println("OperationID: DeleteIncidentTodo")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentTodoCmd)
}
