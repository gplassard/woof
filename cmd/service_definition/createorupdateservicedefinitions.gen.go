package service_definition

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOrUpdateServiceDefinitionsCmd = &cobra.Command{
	Use:   "createorupdateservicedefinitions",
	Short: "Create or update service definition",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/services/definitions")
		fmt.Println("OperationID: CreateOrUpdateServiceDefinitions")
	},
}

func init() {
	Cmd.AddCommand(CreateOrUpdateServiceDefinitionsCmd)
}
