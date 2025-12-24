package service_definition

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetServiceDefinitionCmd = &cobra.Command{
	Use:   "getservicedefinition",
	Short: "Get a single service definition",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/services/definitions/{service_name}")
		fmt.Println("OperationID: GetServiceDefinition")
	},
}

func init() {
	Cmd.AddCommand(GetServiceDefinitionCmd)
}
