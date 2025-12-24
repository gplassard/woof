package case_management_attribute

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCustomAttributeConfigCmd = &cobra.Command{
	Use:   "deletecustomattributeconfig",
	Short: "Delete custom attributes config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cases/types/{case_type_id}/custom_attributes/{custom_attribute_id}")
		fmt.Println("OperationID: DeleteCustomAttributeConfig")
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomAttributeConfigCmd)
}
