package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateReferenceTableCmd = &cobra.Command{
	Use:   "updatereferencetable",
	Short: "Update reference table",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/reference-tables/tables/{id}")
		fmt.Println("OperationID: UpdateReferenceTable")
	},
}

func init() {
	Cmd.AddCommand(UpdateReferenceTableCmd)
}
