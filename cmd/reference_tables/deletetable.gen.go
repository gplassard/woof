package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTableCmd = &cobra.Command{
	Use:   "deletetable",
	Short: "Delete table",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/reference-tables/tables/{id}")
		fmt.Println("OperationID: DeleteTable")
	},
}

func init() {
	Cmd.AddCommand(DeleteTableCmd)
}
