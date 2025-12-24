package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCatalogEntityCmd = &cobra.Command{
	Use:   "deletecatalogentity",
	Short: "Delete a single entity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/catalog/entity/{entity_id}")
		fmt.Println("OperationID: DeleteCatalogEntity")
	},
}

func init() {
	Cmd.AddCommand(DeleteCatalogEntityCmd)
}
