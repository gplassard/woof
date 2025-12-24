package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetProjectCmd = &cobra.Command{
	Use:   "getproject",
	Short: "Get the details of a project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cases/projects/{project_id}")
		fmt.Println("OperationID: GetProject")
	},
}

func init() {
	Cmd.AddCommand(GetProjectCmd)
}
