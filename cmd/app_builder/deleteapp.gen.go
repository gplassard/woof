package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAppCmd = &cobra.Command{
	Use:   "deleteapp",
	Short: "Delete App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/app-builder/apps/{app_id}")
		fmt.Println("OperationID: DeleteApp")
	},
}

func init() {
	Cmd.AddCommand(DeleteAppCmd)
}
