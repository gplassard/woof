package case_management_attribute

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAllCustomAttributesCmd = &cobra.Command{
	Use:   "getallcustomattributes",
	Short: "Get all custom attributes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cases/types/custom_attributes")
		fmt.Println("OperationID: GetAllCustomAttributes")
	},
}

func init() {
	Cmd.AddCommand(GetAllCustomAttributesCmd)
}
