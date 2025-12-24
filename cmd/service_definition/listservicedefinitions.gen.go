package service_definition

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListServiceDefinitionsCmd = &cobra.Command{
	Use:   "listservicedefinitions",
	Short: "Get all service definitions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/services/definitions")
		fmt.Println("OperationID: ListServiceDefinitions")
	},
}

func init() {
	Cmd.AddCommand(ListServiceDefinitionsCmd)
}
