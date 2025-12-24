package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCatalogEntityCmd = &cobra.Command{
	Use:   "listcatalogentity",
	Short: "Get a list of entities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/catalog/entity")
		fmt.Println("OperationID: ListCatalogEntity")
	},
}

func init() {
	Cmd.AddCommand(ListCatalogEntityCmd)
}
