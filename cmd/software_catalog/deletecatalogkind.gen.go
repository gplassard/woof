package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCatalogKindCmd = &cobra.Command{
	Use:   "deletecatalogkind",
	Short: "Delete a single kind",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/catalog/kind/{kind_id}")
		fmt.Println("OperationID: DeleteCatalogKind")
	},
}

func init() {
	Cmd.AddCommand(DeleteCatalogKindCmd)
}
