package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateAttributesCmd = &cobra.Command{
	Use:   "updateattributes",
	Short: "Update case attributes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/attributes")
		fmt.Println("OperationID: UpdateAttributes")
	},
}

func init() {
	Cmd.AddCommand(UpdateAttributesCmd)
}
