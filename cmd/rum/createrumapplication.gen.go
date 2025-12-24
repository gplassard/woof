package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateRUMApplicationCmd = &cobra.Command{
	Use:   "createrumapplication",
	Short: "Create a new RUM application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/rum/applications")
		fmt.Println("OperationID: CreateRUMApplication")
	},
}

func init() {
	Cmd.AddCommand(CreateRUMApplicationCmd)
}
