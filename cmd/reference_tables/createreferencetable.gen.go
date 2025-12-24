package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateReferenceTableCmd = &cobra.Command{
	Use:   "createreferencetable",
	Short: "Create reference table",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/reference-tables/tables")
		fmt.Println("OperationID: CreateReferenceTable")
	},
}

func init() {
	Cmd.AddCommand(CreateReferenceTableCmd)
}
