package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCatalogKindCmd = &cobra.Command{
	Use:   "listcatalogkind",
	Short: "Get a list of entity kinds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/catalog/kind")
		fmt.Println("OperationID: ListCatalogKind")
	},
}

func init() {
	Cmd.AddCommand(ListCatalogKindCmd)
}
