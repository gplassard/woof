package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCatalogRelationCmd = &cobra.Command{
	Use:   "listcatalogrelation",
	Short: "Get a list of entity relations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/catalog/relation")
		fmt.Println("OperationID: ListCatalogRelation")
	},
}

func init() {
	Cmd.AddCommand(ListCatalogRelationCmd)
}
