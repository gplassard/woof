package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var QueryScalarDataCmd = &cobra.Command{
	Use:   "queryscalardata",
	Short: "Query scalar data across multiple products",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/query/scalar")
		fmt.Println("OperationID: QueryScalarData")
	},
}

func init() {
	Cmd.AddCommand(QueryScalarDataCmd)
}
