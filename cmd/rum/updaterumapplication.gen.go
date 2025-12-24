package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateRUMApplicationCmd = &cobra.Command{
	Use:   "updaterumapplication",
	Short: "Update a RUM application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/rum/applications/{id}")
		fmt.Println("OperationID: UpdateRUMApplication")
	},
}

func init() {
	Cmd.AddCommand(UpdateRUMApplicationCmd)
}
