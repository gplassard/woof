package api_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteOpenAPICmd = &cobra.Command{
	Use:   "deleteopenapi",
	Short: "Delete an API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/apicatalog/api/{id}")
		fmt.Println("OperationID: DeleteOpenAPI")
	},
}

func init() {
	Cmd.AddCommand(DeleteOpenAPICmd)
}
