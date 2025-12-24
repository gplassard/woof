package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAppCmd = &cobra.Command{
	Use:   "getapp",
	Short: "Get App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/app-builder/apps/{app_id}")
		fmt.Println("OperationID: GetApp")
	},
}

func init() {
	Cmd.AddCommand(GetAppCmd)
}
