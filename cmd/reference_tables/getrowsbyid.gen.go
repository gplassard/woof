package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRowsByIDCmd = &cobra.Command{
	Use:   "getrowsbyid",
	Short: "Get rows by id",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/reference-tables/tables/{id}/rows")
		fmt.Println("OperationID: GetRowsByID")
	},
}

func init() {
	Cmd.AddCommand(GetRowsByIDCmd)
}
