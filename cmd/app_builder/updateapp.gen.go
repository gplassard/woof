package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateAppCmd = &cobra.Command{
	Use:   "updateapp",
	Short: "Update App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/app-builder/apps/{app_id}")
		fmt.Println("OperationID: UpdateApp")
	},
}

func init() {
	Cmd.AddCommand(UpdateAppCmd)
}
