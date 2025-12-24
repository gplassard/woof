package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpsertRowsCmd = &cobra.Command{
	Use:   "upsertrows",
	Short: "Upsert rows",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/reference-tables/tables/{id}/rows")
		fmt.Println("OperationID: UpsertRows")
	},
}

func init() {
	Cmd.AddCommand(UpsertRowsCmd)
}
