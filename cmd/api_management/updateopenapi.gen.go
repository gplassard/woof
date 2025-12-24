package api_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateOpenAPICmd = &cobra.Command{
	Use:   "updateopenapi",
	Short: "Update an API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/apicatalog/api/{id}/openapi")
		fmt.Println("OperationID: UpdateOpenAPI")
	},
}

func init() {
	Cmd.AddCommand(UpdateOpenAPICmd)
}
