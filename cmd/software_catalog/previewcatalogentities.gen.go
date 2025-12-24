package software_catalog

import (
	"fmt"
	"github.com/spf13/cobra"
)

var PreviewCatalogEntitiesCmd = &cobra.Command{
	Use:   "previewcatalogentities",
	Short: "Preview catalog entities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/catalog/entity/preview")
		fmt.Println("OperationID: PreviewCatalogEntities")
	},
}

func init() {
	Cmd.AddCommand(PreviewCatalogEntitiesCmd)
}
