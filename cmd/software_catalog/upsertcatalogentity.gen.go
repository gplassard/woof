package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpsertCatalogEntityCmd = &cobra.Command{
	Use:   "upsertcatalogentity",
	Short: "Create or update entities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/catalog/entity")
		fmt.Println("OperationID: UpsertCatalogEntity")
	},
}

func init() {
	Cmd.AddCommand(UpsertCatalogEntityCmd)
}
