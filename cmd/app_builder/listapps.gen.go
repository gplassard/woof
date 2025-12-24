package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAppsCmd = &cobra.Command{
	Use:   "listapps",
	Short: "List Apps",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/app-builder/apps")
		fmt.Println("OperationID: ListApps")
	},
}

func init() {
	Cmd.AddCommand(ListAppsCmd)
}
