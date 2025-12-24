package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentTodoCmd = &cobra.Command{
	Use:   "createincidenttodo",
	Short: "Create an incident todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents/{incident_id}/relationships/todos")
		fmt.Println("OperationID: CreateIncidentTodo")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTodoCmd)
}
