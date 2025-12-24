package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAppCmd = &cobra.Command{
	Use:   "createapp",
	Short: "Create App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/app-builder/apps")
		fmt.Println("OperationID: CreateApp")
	},
}

func init() {
	Cmd.AddCommand(CreateAppCmd)
}
