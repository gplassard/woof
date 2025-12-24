package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpsertCatalogKindCmd = &cobra.Command{
	Use:   "upsertcatalogkind",
	Short: "Create or update kinds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/catalog/kind")
		fmt.Println("OperationID: UpsertCatalogKind")
	},
}

func init() {
	Cmd.AddCommand(UpsertCatalogKindCmd)
}
