package case_management_attribute

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCustomAttributeConfigCmd = &cobra.Command{
	Use:   "createcustomattributeconfig",
	Short: "Create custom attribute config for a case type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/types/{case_type_id}/custom_attributes")
		fmt.Println("OperationID: CreateCustomAttributeConfig")
	},
}

func init() {
	Cmd.AddCommand(CreateCustomAttributeConfigCmd)
}
