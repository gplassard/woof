package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteRowsCmd = &cobra.Command{
	Use:   "deleterows",
	Short: "Delete rows",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/reference-tables/tables/{id}/rows")
		fmt.Println("OperationID: DeleteRows")
	},
}

func init() {
	Cmd.AddCommand(DeleteRowsCmd)
}
