package api_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOpenAPICmd = &cobra.Command{
	Use:   "getopenapi",
	Short: "Get an API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/apicatalog/api/{id}/openapi")
		fmt.Println("OperationID: GetOpenAPI")
	},
}

func init() {
	Cmd.AddCommand(GetOpenAPICmd)
}
