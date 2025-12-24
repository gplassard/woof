package api_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOpenAPICmd = &cobra.Command{
	Use:   "createopenapi",
	Short: "Create a new API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/apicatalog/openapi")
		fmt.Println("OperationID: CreateOpenAPI")
	},
}

func init() {
	Cmd.AddCommand(CreateOpenAPICmd)
}
