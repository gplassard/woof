package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCaseCustomAttributeCmd = &cobra.Command{
	Use:   "updatecasecustomattribute",
	Short: "Update case custom attribute",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/custom_attributes/{custom_attribute_key}")
		fmt.Println("OperationID: UpdateCaseCustomAttribute")
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseCustomAttributeCmd)
}
