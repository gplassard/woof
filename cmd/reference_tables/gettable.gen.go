package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTableCmd = &cobra.Command{
	Use:   "gettable",
	Short: "Get table",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/reference-tables/tables/{id}")
		fmt.Println("OperationID: GetTable")
	},
}

func init() {
	Cmd.AddCommand(GetTableCmd)
}
