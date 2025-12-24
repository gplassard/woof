package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCaseCustomAttributeCmd = &cobra.Command{
	Use:   "deletecasecustomattribute",
	Short: "Delete custom attribute from case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cases/{case_id}/custom_attributes/{custom_attribute_key}")
		fmt.Println("OperationID: DeleteCaseCustomAttribute")
	},
}

func init() {
	Cmd.AddCommand(DeleteCaseCustomAttributeCmd)
}
