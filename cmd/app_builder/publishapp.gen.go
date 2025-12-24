package app_builder

import (
	"fmt"
	"github.com/spf13/cobra"
)

var PublishAppCmd = &cobra.Command{
	Use:   "publishapp",
	Short: "Publish App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/app-builder/apps/{app_id}/deployment")
		fmt.Println("OperationID: PublishApp")
	},
}

func init() {
	Cmd.AddCommand(PublishAppCmd)
}
