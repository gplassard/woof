package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAppsCmd = &cobra.Command{
	Use:   "deleteapps",
	Short: "Delete Multiple Apps",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/app-builder/apps")
		fmt.Println("OperationID: DeleteApps")
	},
}

func init() {
	Cmd.AddCommand(DeleteAppsCmd)
}
