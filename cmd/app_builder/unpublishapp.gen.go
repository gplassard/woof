package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UnpublishAppCmd = &cobra.Command{
	Use:   "unpublishapp",
	Short: "Unpublish App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/app-builder/apps/{app_id}/deployment")
		fmt.Println("OperationID: UnpublishApp")
	},
}

func init() {
	Cmd.AddCommand(UnpublishAppCmd)
}
