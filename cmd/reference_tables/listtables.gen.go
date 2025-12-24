package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTablesCmd = &cobra.Command{
	Use:   "listtables",
	Short: "List tables",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/reference-tables/tables")
		fmt.Println("OperationID: ListTables")
	},
}

func init() {
	Cmd.AddCommand(ListTablesCmd)
}
