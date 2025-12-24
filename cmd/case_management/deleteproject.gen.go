package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteProjectCmd = &cobra.Command{
	Use:   "deleteproject",
	Short: "Remove a project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cases/projects/{project_id}")
		fmt.Println("OperationID: DeleteProject")
	},
}

func init() {
	Cmd.AddCommand(DeleteProjectCmd)
}
