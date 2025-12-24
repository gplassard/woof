package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteRUMApplicationCmd = &cobra.Command{
	Use:   "deleterumapplication",
	Short: "Delete a RUM application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/rum/applications/{id}")
		fmt.Println("OperationID: DeleteRUMApplication")
	},
}

func init() {
	Cmd.AddCommand(DeleteRUMApplicationCmd)
}
