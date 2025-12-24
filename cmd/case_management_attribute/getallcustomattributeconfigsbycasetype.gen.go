package case_management_attribute

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAllCustomAttributeConfigsByCaseTypeCmd = &cobra.Command{
	Use:   "getallcustomattributeconfigsbycasetype",
	Short: "Get all custom attributes config of case type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cases/types/{case_type_id}/custom_attributes")
		fmt.Println("OperationID: GetAllCustomAttributeConfigsByCaseType")
	},
}

func init() {
	Cmd.AddCommand(GetAllCustomAttributeConfigsByCaseTypeCmd)
}
