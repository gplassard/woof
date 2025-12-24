package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetProjectsCmd = &cobra.Command{
	Use:   "getprojects",
	Short: "Get all projects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cases/projects")
		fmt.Println("OperationID: GetProjects")
	},
}

func init() {
	Cmd.AddCommand(GetProjectsCmd)
}
