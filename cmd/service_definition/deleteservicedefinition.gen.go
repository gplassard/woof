package service_definition

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteServiceDefinitionCmd = &cobra.Command{
	Use:   "deleteservicedefinition",
	Short: "Delete a single service definition",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/services/definitions/{service_name}")
		fmt.Println("OperationID: DeleteServiceDefinition")
	},
}

func init() {
	Cmd.AddCommand(DeleteServiceDefinitionCmd)
}
