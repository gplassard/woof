package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateProjectCmd = &cobra.Command{
	Use:   "createproject",
	Short: "Create a project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/projects")
		fmt.Println("OperationID: CreateProject")
	},
}

func init() {
	Cmd.AddCommand(CreateProjectCmd)
}
